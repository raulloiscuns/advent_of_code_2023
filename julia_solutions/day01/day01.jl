module Day01

export run

# Function for the problem
function run(input::String)
    lines = readlines(input)
    part_1 = part1(lines)
    part_2 = part2(lines)
    return [part_1, part_2]
end

# Function for part1 of the problem
function part1(lines::Vector{String})
    value = 0
    for line in lines
        numbers = []
        for charac in line
            try
                number = parse(Int, charac)
                append!(numbers, number)
            catch
                continue
            end
        end
        value +=  10*numbers[1] + numbers[end]
    end
    return value
end

function part2(lines::Vector{String})
    value = 0
    for line in lines
        numbers = []
        for (index,charac) in enumerate(line)
            try
                number = parse(Int, charac)
                append!(numbers, number)
            catch
                if startswith(line[index:end], "one")
                    append!(numbers, 1)
                elseif startswith(line[index:end], "two")
                    append!(numbers, 2)
                elseif startswith(line[index:end], "three")
                    append!(numbers, 3)
                elseif startswith(line[index:end], "four")
                    append!(numbers, 4)
                elseif startswith(line[index:end], "five")
                    append!(numbers, 5)
                elseif startswith(line[index:end], "six")
                    append!(numbers, 6)
                elseif startswith(line[index:end], "seven")
                    append!(numbers, 7)
                elseif startswith(line[index:end], "eight")
                    append!(numbers, 8)
                elseif startswith(line[index:end], "nine")
                    append!(numbers, 9)
                end
            end
        end
        value +=  10*numbers[1] + numbers[end]
    end
    return value
end

end