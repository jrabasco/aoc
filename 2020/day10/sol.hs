-- {-# LANGUAGE BangPatterns #-}
import System.IO
import Data.List (sort)
import Data.Function (fix)
import qualified Data.Map as Map

main = do
    contents <- readFile "input.txt"
    let joltages =  readInts (lines contents)
    print(part1 joltages)
    print(part2 joltages)


readInts :: [String] -> [Int]
readInts = map read

findDiffs :: [Int] -> Int -> [Int]
findDiffs joltages start
    | null joltages = []
    | otherwise = ((head joltages) - start):(findDiffs (tail joltages) (head joltages))

collectDiffs :: [Int] -> (Int, Int) -> (Int, Int)
collectDiffs diffs (diff1, diff3)
    | null diffs = (diff1, diff3)
    | (head diffs) == 1 = collectDiffs (tail diffs) ((diff1+1), diff3)
    | (head diffs) == 3 = collectDiffs (tail diffs) (diff1, (diff3+1))
    | otherwise = collectDiffs (tail diffs) (diff1, diff3)

countVars :: [Int] -> Map.Map Int Int -> Int -> Int
countVars jolts memo end
    | end == 0 = 1
    | otherwise =  case (Map.lookup end memo) of
        Just count -> count
        Nothing -> let (_, count) = foldl (\(mem, acc) c ->
                             case (Map.lookup c mem) of
                                 Just count -> (mem, acc + count)
                                 Nothing ->
                                     if ((jolts!!end)-(jolts!!c) <= 3)
                                     then let cCount = countVars jolts mem c
                                              newMem = Map.insert c cCount mem
                                          in (newMem, acc + cCount)
                                     else (mem, acc)) (memo, 0) [0..end-1]
                   in count

pathCounts :: [Int] -> Int -> [Int] -> [Int]
pathCounts rjolts from paths =
    let start = rjolts!!from
        newPaths = [ paths!!(n-1) | n <- [1..3], from - n >= 0, (rjolts!!(from-n) - start) <= 3 ]
        newCount = sum newPaths
    in newCount : paths

part1 :: [Int] -> Int
part1 joltages = let diffs = (findDiffs (sort joltages) 0)
                     (diff1, diff3) = collectDiffs diffs (0,1)
                 in diff1*diff3

part2 :: [Int] -> Int
part2 joltages = let mJolt = (maximum joltages) + 3
                     jolts = sort joltages
                     allJolts = 0:jolts ++ [mJolt]
                     paths 0 = [1]
                     paths i = pathCounts (reverse allJolts) i (paths (i-1))
                 in  head $ paths ((length allJolts) - 1)
