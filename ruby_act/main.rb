def is_palindrome?(str)
    proc_str = str.downcase.gsub(/[^a-zA-Z0-9]/, '')

    reverse_str = proc_str.reverse

    proc_str == reverse_str
end

def permutations(str)
    str.chars.to_a.permutation.uniq.length()
end

def int_to_roman(int)
    roman_values = {
        1000 => "M", 900 => "CM", 500 => "D",
        400 => "CD", 100 => "C", 90 => "XC",
        50 => "L", 40 => "XL", 10 => "X",
        9 => "IX", 5 => "V", 4 => "IV", 
        1 => "I"
    }

    roman = ""
    roman_values.each do |key, value|
        while int >= key
            roman << value
            int -= key
        end
    end
    roman
end

def is_armstrong?(int)
    int_arr = int.digits
    sum = 0
    int_arr.each do |i|
        sum += i**int_arr.length()
    end
    sum == int
end

def main()
    puts "If the input is a combination of numbers and letters, it is treated as a string."
    print "Enter a string or integer: "
    input = gets.chomp

    begin
        int_val = Integer(input)
        puts ""
        puts "#{int_val} to roman numeral: #{int_to_roman(int_val)}"
        if is_armstrong?(int_val)
            puts "#{int_val} is an armstrong number!"
        else
            puts "#{int_val} is not an armstrong number."
        end
    rescue ArgumentError
        puts ""
        if is_palindrome?(input)
            puts "#{input} is a palindrome and has #{permutations(input)} permutations!"
        else 
            puts "#{input} is not a palindrome and has #{permutations(input)} permutations."
        end
    end
end

main()