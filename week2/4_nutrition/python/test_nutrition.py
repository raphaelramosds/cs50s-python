from nutrition import nutrition


def test_apple():
    assert nutrition("Apple") == "Calories: 130"


def test_avocado():
    assert nutrition("Avocado") == "Calories: 50"


def test_sweet_cherries():
    assert nutrition("Sweet Cherries") == "Calories: 100"
