from meal import meal


def test_seven_oclock():
    assert meal("7:00") == "breakfast time"


def test_seven_thirty():
    assert meal("7:30") == "breakfast time"


def test_twelve_fourty_two():
    assert meal("12:42") == "lunch time"


def test_eighteen_thirty_two():
    assert meal("18:32") == "dinner time"


def test_nothing_time():
    assert meal("11:11") == ""
