#!/usr/bin/python3.8
from dataclasses import dataclass
from typing import List, Tuple, Optional, TypeVar, Callable

@dataclass
class Packet:
    version: int
    type_id: int

@dataclass
class Literal(Packet):
    value: int


@dataclass
class Operator(Packet):
    length_type: int
    sub_packets: List[Packet]


def product(in_list: List[int]) -> int:
    acc = 1
    for elm in in_list:
        acc *= elm
    return acc

def gt(in_list: List[int]) -> int:
    if in_list[0] > in_list[1]:
        return 1
    return 0

def lt(in_list: List[int]) -> int:
    if in_list[0] < in_list[1]:
        return 1
    return 0

def eq(in_list: List[int]) -> int:
    if in_list[0] == in_list[1]:
        return 1
    return 0

OPERATORS = {
    0 : sum,
    1 : product,
    2 : min,
    3 : max,
    5 : gt,
    6 : lt,
    7 : eq
}


def get_bit_string(transmisison: str) -> str:
    b_size = len(transmission)*4
    return bin(int(transmission, 16))[2:].zfill(b_size)

# Parses a literal value and returns the value and the number of bits consumed
def parse_literal(bitstring: str, start: int) -> Tuple[int, int]:
    lead = '1'
    value_str = ''
    it = start
    while lead != '0':
        chunk = bitstring[it:it+5]
        lead = chunk[0]
        value_str += chunk[1:]
        it += 5
    return int(value_str, 2), it - start

# Parses an operator packet, returns the length type, list of sub_packets
# and the number of bits consumed
def parse_operator(bitstring: str, start: int)-> Tuple[int, List[Packet], int]:
    length_type = int(bitstring[start])
    it = start + 1
    sub_bits_length = -1
    nb_sub_packets = -1
    if length_type == 0:
        sub_bits_length = int(bitstring[it:it+15], 2)
        it += 15
    else:
        nb_sub_packets = int(bitstring[it:it+11], 2)
        it += 11

    sub_bits_read = 0
    sub_packets = []
    while sub_bits_read < sub_bits_length or len(sub_packets) < nb_sub_packets:
        packet, consumed = parse_packet(bitstring, it)
        it += consumed
        sub_bits_read += consumed
        sub_packets.append(packet)
    return length_type, sub_packets, it - start

# Parses a packet and returns the packet and how many bits were consumed
def parse_packet(bitstring: str, start: int) -> Tuple[Packet, int]:
    version = int(bitstring[start:start+3], 2)
    type_id = int(bitstring[start+3:start+6], 2)

    if type_id == 4:
        value, consumed = parse_literal(bitstring, start+6)
        return Literal(version, type_id, value), consumed + 6
    else:
        l_type, packets, consumed = parse_operator(bitstring, start+6)
        return Operator(version, type_id, l_type, packets), consumed + 6

# Parse a whole transmission
def parse(transmission: str) -> Packet:
    bitstring = get_bit_string(transmission)
    packet, _ = parse_packet(bitstring, 0)
    return packet

# Pre-order traversal which returns a list of values extracted using the
# extractor function
T = TypeVar('T')
def pre_order_traversal(root: Packet,
                        extractor: Callable[[Packet], T]) -> List[T]:
    res = [extractor(root)]
    # packet is an operator
    if root.type_id != 4:
        for p in root.sub_packets:
            res += pre_order_traversal(p, extractor)
    return res


def sum_versions(packet: Packet) -> int:
    return sum(pre_order_traversal(packet, lambda p: p.version))

def evaluate(root: Packet) -> int:
    # literal
    if root.type_id == 4:
        return root.value
    params = [ evaluate(p) for p in root.sub_packets ]
    op = OPERATORS[root.type_id]
    return op(params)

for i in range(1, 8):
    transmission = open(f'p1input{i}.txt').read().strip()
    packet = parse(transmission)
    print(f'Transmission: {transmission}, sum of versions: {sum_versions(packet)}')

for i in range(1, 9):
    transmission = open(f'p2input{i}.txt').read().strip()
    packet = parse(transmission)
    print(f'Transmission: {transmission}, evaluates to: {evaluate(packet)}')

transmission = open(f'input.txt').read().strip()
packet = parse(transmission)
print(f'Part 1: {sum_versions(packet)}')
print(f'Part 2: {evaluate(packet)}')
