#!/usr/bin/python3.7

digits = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

def valid_height(hgt):
    if "cm" in hgt:
        height = hgt.split("cm")[0]
        return 150 <= int(height) <= 193
    if "in" in hgt:
        height = hgt.split("in")[0]
        return 59 <= int(height) <= 76
    return False

def valid_hair_color(hcl):
    if len(hcl) == 0 or hcl[0] != "#":
        return False

    chars = hcl[1:]
    if len(chars) != 6:
        return False
    valid_chars = digits | {'a', 'b', 'c', 'd', 'e', 'f'}
    return all(c in valid_chars for c in chars)

validation = {
        "byr": lambda byr: 1920 <= int(byr) <= 2002,
        "iyr": lambda iyr: 2010 <= int(iyr) <= 2020,
        "eyr": lambda eyr: 2020 <= int(eyr) <= 2030,
        "hgt": valid_height,
        "hcl": valid_hair_color,
        "ecl": lambda ecl: ecl in {'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'},
        "pid": lambda pid: len(pid) == 9 and all(c in digits for c in pid)
}

lines = [ l.strip() for l in open('input.txt').readlines() ]
passports = [{}]

for line in lines:
    if line == '':
        passports.append({})
        continue

    entries = line.split(" ")
    for entry in entries:
        key, value = entry.split(":")
        passports[-1][key] = value

valid_passports = 0
for passport in passports:
    if all(field in passport and validation[field](passport[field]) for field in validation):
        valid_passports += 1

print(valid_passports)
