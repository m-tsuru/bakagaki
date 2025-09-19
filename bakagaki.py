import random


def random_subsequence(s: str) -> str:
    return "".join(ch for ch in s if bool(random.getrandbits(1)))


s = "バーガーキング"
while True:
    print(random_subsequence(s))
