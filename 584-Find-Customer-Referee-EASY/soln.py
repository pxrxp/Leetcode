import pandas as pd

def find_customer_referee(customer: pd.DataFrame) -> pd.DataFrame:
    null_mask = customer['referee_id'].isna()
    not_two_mask = customer['referee_id'] != 2

    return customer.loc[null_mask | not_two_mask, ['name']]
