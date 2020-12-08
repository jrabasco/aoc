import System.IO

main = do
    contents <- readFile "input.txt"
    let instructions = lines contents
    print(part1 instructions)
    print(part2 instructions)

readInt :: String -> Int
readInt ('+':rest) = read rest
readInt any = read any

simulate :: [String] -> [Int] -> Int -> Int -> (Int, Bool)
simulate instructions visited ip acc
    | elem ip visited = (acc, False)
    | ip >= length instructions = (acc, True)
    | otherwise = case words(instructions!!ip) of
        "nop":nb:[] -> simulate instructions (ip:visited) (ip + 1) acc
        "acc":nb:[] -> simulate instructions (ip:visited) (ip + 1) (acc + (readInt nb))
        "jmp":nb:[] -> simulate instructions (ip:visited) (ip + (readInt nb)) acc

swap :: Int -> [String] -> [String]
swap ip instructions =
    case words(instructions!!ip) of
        "nop":nb:[] -> (take ip instructions) ++ (("jmp " ++ nb) : (drop (ip+1) instructions))
        "jmp":nb:[] -> (take ip instructions) ++ (("nop " ++ nb) : (drop (ip+1) instructions))
        any -> instructions

trySwap :: Int -> Int -> [String] -> (Int, Bool)
trySwap ip acc instructions
    | ip >= length instructions = (acc, False)
    | otherwise =
        case (simulate (swap ip instructions) [] 0 0) of
            (acc, True) -> (acc, True)
            (acc, False) -> trySwap (ip+1) acc instructions

part1 :: [String] -> Int
part1 instructions =
  let (res,_) = simulate instructions [] 0 0
  in res

part2 :: [String] -> (Int, Bool)
part2 instructions = trySwap 0 0 instructions

