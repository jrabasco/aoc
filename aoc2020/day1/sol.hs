import System.IO

main = do
    contents <- readFile "input.txt"
    let numbers = readInt . lines $ contents
    print(part1 numbers)
    print(part2 numbers)


readInt :: [String] -> [Int]
readInt = map read

part1 :: [Int] -> Int
part1 numbers
  | length numbers <= 2 = 0
  | otherwise = fst res * snd res
                where res = head $ findSums2 2020 numbers

part2 :: [Int] -> Int
part2 numbers
  | length numbers <= 3 = 0
  | otherwise = let (x,y,z) = head $ findSums3 2020 numbers
                in x*y*z


findSums2 :: Int -> [Int] -> [(Int, Int)]
findSums2 sm numbers = map (\x -> (x, sm-x)) (filter (\x -> elem (sm - x) numbers) numbers)

findSums3 :: Int -> [Int] -> [(Int, Int, Int)]
findSums3 sm numbers = [(x, y, z) | x <- numbers, y <- numbers, z <- numbers, x+y+z == sm]
