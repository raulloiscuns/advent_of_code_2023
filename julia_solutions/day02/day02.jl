module Day02

export run

# Function for the problem
function run(input::String)
    data = readinput(input)
    part_1 = part1(data)
    part_2 = part2(data)
    return [part_1, part_2]
end

# Function for read the input
function readinput(input::String)
    lines = readlines(input)
    game_dic = Dict()
    index_regex = r"Game (\d+):" 
    red_regex = r"(\d+) red"
    green_regex = r"(\d+) green"
    blue_regex = r"(\d+) blue"
    for line in lines
        index = match(index_regex, line).captures[1]
        red = parse.(Int, [m.captures[1] for m in eachmatch(red_regex, line)])
        green = parse.(Int, [m.captures[1] for m in eachmatch(green_regex, line)])
        blue = parse.(Int, [m.captures[1] for m in eachmatch(blue_regex, line)])
        game_dic[index]=[red, green, blue]
    end
    return game_dic
end

# Function for part1 of the problem
function part1(data)
    sum = 0
    for (key, value) in data
        if prod(value[1].<=12) && prod(value[2].<=13) && prod(value[3].<=14)
            sum += parse(Int, key)
        end
    end
    return sum
end

function part2(data)
    sum = 0
    for (_, value) in data
        sum += maximum(value[1])*maximum(value[2])*maximum(value[3])
    end
    return sum
end

end