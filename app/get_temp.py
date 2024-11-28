"""
Module for getting the average temperature of all sensebox sensors.
"""

from datetime import datetime, timezone, timedelta
import requests


def get_sensor_data():
    """
    Get the sensor data from the sensebox API.
    """
    current_time = datetime.now(timezone.utc)
    one_hour_ago = current_time - timedelta(hours=1)
    current_time = format_time(current_time)
    one_hour_ago = format_time(one_hour_ago)
    api_url = "https://api.opensensemap.org/boxes"

    params = {
        "phenomenon": "temperature",
        "date": current_time,
    }

    response = requests.get(api_url, params=params, timeout=30)

    if response.status_code != 200:
        raise requests.exceptions.RequestException(
            f"failed to get sensor data: {response.json()}"
        )

    json = response.json()
    return json


def format_time(time):
    """
    Format the time to the correct format.
    """
    return time.strftime("%Y-%m-%dT%H:%M:%SZ")


def filter_sensor_data(sensor_data):
    """
    Filter the sensor data to only include temperature values that are less than an hour old.
    """
    temperature_values = []
    for box in sensor_data:
        if "sensors" not in box.keys():
            continue
        for sensor in box["sensors"]:
            if "temperature" == sensor["title"]:
                temperature_values.append(float(sensor["lastMeasurement"]["value"]))
            elif "Temperatur" in sensor["title"]:
                temperature_values.append(float(sensor["lastMeasurement"]["value"]))
    return temperature_values


def calculate_average_temperature(sensor_data):
    """
    Calculate the average temperature of all sensebox sensors.
    """

    return round(sum(sensor_data) / len(sensor_data), 2)


def get_average_tempeature() -> float:
    """
    Get the average temperature of all sensebox sensors.
    """
    try:
        data = get_sensor_data()
        temps = filter_sensor_data(data)
        return calculate_average_temperature(temps)
    except Exception as e:
        raise e
