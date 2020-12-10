{-# LANGUAGE BangPatterns #-}
import System.IO
import Data.List
import Data.Function (fix)

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

data Tree a = Tree (Tree a) a (Tree a)
instance Functor Tree where
    fmap f (Tree l m r) = Tree (fmap f l) (f m) (fmap f r)

index :: Tree a -> Int -> a
index (Tree _ m _) 0 = m
index (Tree l _ r) n = case (n - 1) `divMod` 2 of
    (q, 0) -> index l q
    (q, 1) -> index r q

nats :: Tree Int
nats = go 0 1
    where
        go !n !s = Tree (go l s') n (go r s')
            where
                l = n + s
                r = l + s
                s' = s * 2


toList :: Tree a -> [a]
toList as = map (index as) [0..]


countVars :: ([Int] -> Int -> Int) -> [Int] -> Int -> Int
countVars f jolts end
    | end == 0 = 1
    | otherwise = foldl (\cnt c -> if ((jolts!!end)-(jolts!!c) <= 3) then (cnt + (f jolts c)) else cnt) 0 [0..end-1]

memoize :: ([Int] -> Int -> a) -> [Int] -> (Int -> a)
memoize f jolts = index (fmap (f jolts) nats)

countVarsMemo :: [Int] -> Int -> Int
countVarsMemo = fix (memoize . countVars)

part1 :: [Int] -> Int
part1 joltages = let diffs = (findDiffs (sort joltages) 0)
                     (diff1, diff3) = collectDiffs diffs (0,1)
                 in diff1*diff3

part2 :: [Int] -> Int
part2 joltages = let jolts = sort joltages
                     allJolts = 0:jolts
                 in countVarsMemo allJolts ((length allJolts)-1)
