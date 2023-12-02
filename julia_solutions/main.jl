using Printf

# Including solutions 
include("day01/day01.jl")

if isempty(ARGS)
    global day = 1
else
    try
        global day = parse(Int, ARGS[1])
    catch
        error(ARGS[1]*" is not a valid day.")
    end
end

if day>25
    error("Day "*string(day)*" is not a valid day\n")
end

@printf("Running day %02d\n",day)

if day == 1
    part_1, part_2 = Day01.run("day01/input.txt")
    println("-> Part 1: "*string(part_1))
    println("-> Part 2: "*string(part_2))
else
    @printf("Day %02d has not been yet solved\n",day)
end