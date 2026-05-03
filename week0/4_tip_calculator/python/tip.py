import re


class InvalidDollarStr(ValueError):
    def __init__(self, txt: str):
        super().__init__(f"dollar: {txt} should be formatted as $##.##")


class InvalidPercentStr(ValueError):
    def __init__(self, txt: str):
        super().__init__(f"percent: {txt} should be formatted as ##%")


def dollars_to_float(dollars: str) -> float:
    if not re.fullmatch(r"\$\d{1,3}(\.\d{1,2})?", dollars):
        raise InvalidDollarStr(dollars)
    return float(dollars[1:])


def percent_to_float(percent: str) -> float:
    if not re.fullmatch(r"\d{1,2}%", percent):
        raise InvalidPercentStr(percent)
    return float(percent[:-1])


def tip_calculator(dollars: str, percent: str) -> str:
    amount = dollars_to_float(dollars)
    pct = percent_to_float(percent)
    tip = amount * (pct / 100)
    return f"Leave ${tip:.2f}"
