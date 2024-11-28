import google.generativeai as genai
import os

# Configure the API with your API key
genai.configure(api_key="")

# Initialize the GenerativeModel with the model ID
model = genai.GenerativeModel("gemini-exp-1121")

# Prepare the prompt for content generation
prompt = """
Users

50,000 users distributed across the year.
Attributes like registration date, type, activity frequency, and location.
User activity simulated based on assigned activity frequency.
Warehouses

20 warehouses distributed geographically.
Attributes like location, capacity, and activation date.
Inventory

500,000 items spread across the warehouses.
Attributes like item category, price, stock quantity, and associated warehouse.
Traffic Events

Includes item views, purchases, returns, and management actions.
Generated based on quarterly and regional trends.
Traffic Patterns
Seasonality:

Q1: Moderate.
Q2: Slight growth.
Q3: Steady with occasional spikes.
Q4: Peak due to holiday shopping.
Event Distribution:

Views: 50%.
Purchases: 30%.
Returns: 10%.
Stock/Management actions: 10%.
Regional Factors:

Urban warehouses have 2x traffic compared to rural ones.
Time Granularity

Predict Traffic for Q5:

Use regression models or seasonal forecasting to estimate future trends.
Generate predictions per event type, per warehouse.
Pod Scaling Recommendations:

Analyze resource usage from the emulated data.
Suggest pod counts per service based on predicted traffic.

Users, warehouses, inventory, and traffic events for four quarters.
Quarter 5 Predictions:

Predicted traffic counts.
Confidence intervals.
Scaling Recommendations:

Predict future pod counts with justifications.

Sample Data Preview
Users
user_id    registration_date    user_type    activity_frequency    location
U000001    2023-04-13         admin        weekly                 West
U000002    2023-12-15         normal       weekly                 West
U000003    2023-09-28         normal       monthly                South
U000004    2023-04-17         normal       weekly                 East
U000005    2023-03-13         normal       daily                  South
Warehouses
warehouse_id    location    capacity    active_since
W001    West    15148    2023-06-02
W002    South   15050    2021-08-16
W003    West    15898    2022-05-17
W004    North   45766    2020-05-28
W005    North   14152    2020-01-24
Inventory
inventory_id    item_name    category    price    stock_quantity    warehouse_id
I000001    Item_1    Books    293.38    349    W014
I000002    Item_2    Electronics   34.74    226    W020
I000003    Item_3    Books    176.91    88     W009
I000004    Item_4    Food     607.35    530    W017
I000005    Item_5    Electronics   193.27   195    W008
Next Steps

Predict Quarter 5 Data
Use the generated data to estimate traffic for Q5 using forecasting models.

write output which includes the suggestions and predictions (and no more filler words) in not more than 10 words.
"""

try:
    response = model.generate_content(prompt)
    print(response.text)
except Exception as e:
    print("Error:", e)
