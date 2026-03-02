import pandas as pd

def find_products(products: pd.DataFrame) -> pd.DataFrame:
    mask = (products['low_fats'] == 'Y') & (products['recyclable'] == 'Y')
    return products.loc[mask, ['product_id']]
    
    # df = products[
    #     (products['low_fats'] == 'Y') &
    #     (products['recyclable'] == 'Y')
    # ]
    # return df[['product_id']]
