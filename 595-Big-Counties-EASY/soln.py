import pandas as pd

def big_countries(world: pd.DataFrame) -> pd.DataFrame:
    big_area_mask = world['area'] >= 3000000
    big_population_mask = world['population'] >= 25000000

    cols = ['name', 'population', 'area']
    return world.loc[big_area_mask | big_population_mask, cols]
