COKE_PRICE = 50
VALID_COINS = [5, 10, 25]

# Global variables
# NOTE: DO NOT USE coke_machine OR coke_machine_static ON THE SAME PROCESS!!
#       This is because they share the same global variables, and will interfere with each other.
amount = 0
change = 0

def process(coin: str) -> bool:
    global amount
    global change

    if not coin.isnumeric():
        print("Please, input a number")
        return False

    coin = int(coin)
    if not coin in VALID_COINS:
        print(
            "Please, input a valid coin ({})".format(
                ", ".join(str(i) for i in VALID_COINS)
            )
        )
        return False

    amount = amount + coin
    if amount >= COKE_PRICE:
        change = amount - COKE_PRICE

    return True


def coke_machine_static(inputs: list) -> dict:
    """
    Testable version of coke_machine which does not depend upon stdin
    """
    for coin in inputs:
        if not process(str(coin)):
            continue

    return {"total_inserted": amount, "coke_price": COKE_PRICE, "change": change}


def coke_machine() -> None:
    while amount < COKE_PRICE:
        print("Amount Due: {}".format(COKE_PRICE - amount))
        coin = input("Insert Coin: ")
        if not process(coin):
            continue

    print("Total inserted: {}".format(amount))
    print("Coke price: {}".format(COKE_PRICE))
    print("Change Owed: {}".format(change))
