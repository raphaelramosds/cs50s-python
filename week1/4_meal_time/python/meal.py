def convert(time_str: str) -> float:
    hh, mm = time_str.strip().split(":")
    return float(hh) + float(mm) / 60


def meal(time_str: str) -> str:
    t = convert(time_str)
    if t >= 7.0 and t <= 8.0:
        return "breakfast time"
    elif t >= 12.0 and t <= 13.0:
        return "lunch time"
    elif t >= 18.0 and t <= 19.0:
        return "dinner time"
    else:
        return ""
