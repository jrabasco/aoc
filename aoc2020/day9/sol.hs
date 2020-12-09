import System.IO

main = do
    contents <- readFile "input.txt"
    let numbers = map readInt $ lines contents
    print(part1 numbers)
    print(part2 numbers)


readInt :: String -> Int
readInt = read

isSumFrom :: Int -> [Int] -> Bool
isSumFrom goal numbers
    | null numbers = False
    | elem (goal - (head numbers) ) (tail numbers) = True
    | otherwise = isSumFrom goal (tail numbers)

findMistake :: [Int] -> [Int] -> Int -> Int
findMistake numbers preamble pSize
    | null numbers = 0
    | length preamble < pSize = findMistake (tail numbers) (preamble ++ [head numbers]) pSize
    | not (isSumFrom (head numbers) preamble) = head numbers
    | otherwise = findMistake (tail numbers) ( (tail preamble) ++ [head numbers]) pSize

findSumTo :: [Int] -> Int -> [Int] -> Int -> [Int]
findSumTo numbers goal found curSum
   | null numbers = []
   | (head numbers + curSum) == goal = found ++ [head numbers]
   | (head numbers + curSum) < goal = findSumTo (tail numbers) goal (found ++ [head numbers]) (curSum + (head numbers))
   | (head numbers + curSum) > goal = findSumTo numbers goal (tail found) (curSum - (head found))

part1 :: [Int] -> Int
part1 numbers = findMistake numbers [] 25

part2 :: [Int] -> Int
part2 numbers =
  let goal = part1 numbers
      found = findSumTo numbers goal [] 0
  in (minimum found) + (maximum found)
