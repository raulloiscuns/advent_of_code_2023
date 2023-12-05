module Day04

export run

# Function for the problem
function run(input::String)
    card_data = readinput(input)
    part_1 = part1(card_data)
    part_2 = part2(card_data)
    return [part_1, part_2]
end

# Function for part1 of the problem
function part1(card_data::Vector)
    points = 0
    for card in card_data
        count = [1 for number in card[2] if number in card[1]] |> sum
        points = count != 0 ? points + 2^(count-1) : points
    end
    return round(Int, points)
end

# Function for part2 of the problem
function part2(card_data::Vector)
    number_of_cards = ones(length(card_data))
    for (i, card) in enumerate(card_data)
        for _ in 1:number_of_cards[i]
            count = [1 for number in card[2] if number in card[1]] |> sum
            for j in i+1:i+count
                number_of_cards[j] += 1
            end
        end
    end
    return round(Int, sum(number_of_cards))
end

# Function for reading the input of the problem
function readinput(input::String)
    lines = readlines(input)
    card_data = []
    for line in lines
        numbers = split(split(line, ':')[2], '|')
        winning_numbers = parse.(Int, SubString.(numbers[1], findall(r"(\d+)", numbers[1])))
        player_numbers = parse.(Int, SubString.(numbers[2], findall(r"(\d+)", numbers[2])))
        push!(card_data, [winning_numbers, player_numbers])
    end
    return card_data
end

end