module Day03

export run

# Function for the problem
function run(input::String)
    part_1 = part1(input)
    part_2 = part2(input)
    return [part_1, part_2]
end

# Function for part1 of the problem
function part1(input::String)
    NOT_SYMBOLS = ['.', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0']
    DIRECTIONS = [[1,0], [-1, 0], [0,1], [0,-1], [1,1], [1,-1], [-1,1], [-1,-1]]
    lines = readlines(input) #Reading input data
    count = 0
    for (i, line) in enumerate(lines)
        ranges =  findall(r"(\d+)", line)
        for range in ranges
            surrounded = false
            for j in range
                surrounded = [!(lines[dir[1]][dir[2]] in NOT_SYMBOLS) for dir in DIRECTIONS .+ [[i,j]] if 1 <= dir[1] < length(lines) && 1 <= dir[2] < length(line)] |> any
                if surrounded == true
                    count += parse(Int, prod(line[range]))
                    break
                end
            end
        end
    end
    return count
end

# Function for part2 of the problem
function part2(input::String)
    DIRECTIONS = [[1,0], [-1, 0], [0,1], [0,-1], [1,1], [1,-1], [-1,1], [-1,-1]]
    lines = readlines(input) #Reading input data
    dic = Dict()
    for (i, line) in enumerate(lines)  
        ranges =  findall(r"(\d+)", line)
        for range in ranges
            surrounded = false 
            for j in range
                for pos in DIRECTIONS .+ [[i,j]]
                    if 1 <= pos[1] < length(lines) && 1 <= pos[2] < length(line) && lines[pos[1]][pos[2]]=='*'
                        try
                            append!(dic[(pos[1],pos[2])],parse(Int, prod(line[range])))
                            surrounded = true
                        catch
                            dic[(pos[1],pos[2])]=[parse(Int, prod(line[range]))]
                            surrounded = true
                        end
                        break
                    end
                end
                surrounded && break 
            end
        end
    end
    return [value[1]*value[2] for (_, value) in dic if length(value)==2] |> sum
end

end