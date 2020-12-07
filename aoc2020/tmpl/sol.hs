import System.IO

main = do
    contents <- readFile "input.txt"
    print(lines contents)
