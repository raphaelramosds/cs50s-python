import pytest
from tip import dollars_to_float, percent_to_float, tip_calculator, InvalidDollarStr, InvalidPercentStr

def test_invalid_dollar_string():
    with pytest.raises(InvalidDollarStr):
        dollars_to_float("$ 7.50")

def test_invalid_percent_string():
    with pytest.raises(InvalidPercentStr):
        percent_to_float("5 %")

def test_5_dollars_to_float():
    assert dollars_to_float("$5") == 5.0

def test_5_percent_to_float():
    assert percent_to_float("5%") == 5.0

def test_50_dollars_15_percent():
    assert tip_calculator("$50", "15%") == "Leave $7.50"

def test_100_dollars_18_percent():
    assert tip_calculator("$100.00", "18%") == "Leave $18.00"

def test_15_dollars_25_percent():
    assert tip_calculator("$15.00", "25%") == "Leave $3.75"
