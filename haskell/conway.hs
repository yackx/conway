import qualified Data.Set as Set

type Cell = (Int, Int)
type Cells = Set.Set Cell

-- Neighbours of a cell, living or dead
neighbours :: Cell -> Cells
neighbours (x, y) = Set.fromList
    [ (x+dx, y+dy)
    | dx <- [-1..1]
    , dy <- [-1..1]
    , (dx, dy) /= (0, 0) -- Exclude the cell itself
    ]

countLivingNeighbours :: Cell -> Cells -> Int
countLivingNeighbours cell cells = Set.size $ Set.intersection (neighbours cell) cells

isAlive :: Cell -> Cells -> Bool
isAlive cell cells = Set.member cell cells

-- Candidate cells are all living cells and their neighbours
candidates :: Cells -> Cells
candidates grid = Set.unions [grid, Set.unions (Set.map neighbours grid)]

-- Apply game's rules
conway :: Cells -> Cells
conway grid = Set.filter (willLive grid) (candidates grid)
  where
    willLive grid cell
      | isAlive cell grid = n == 2 || n == 3
      | otherwise         = n == 3
      where n = countLivingNeighbours cell grid

displayGrid :: Cells -> IO ()
displayGrid cells = putStrLn $ show (Set.toList cells)

initialState :: Cells
initialState = Set.fromList [(0, 0), (1, 0), (2, 0), (1, 1), (2, 1)]

main :: IO ()
main = do
  displayGrid initialState
  runSimulation initialState
  where
    runSimulation grid 
      | Set.null grid = return ()
      | otherwise = do
          let next = conway grid
          displayGrid next
          runSimulation next
