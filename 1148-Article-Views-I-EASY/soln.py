import pandas as pd

def self_views(views: pd.DataFrame) -> pd.DataFrame:
    mask = views["author_id"] == views["viewer_id"]
    cols = ["author_id"]
    
    result = views.loc[mask, cols]
    result = result.drop_duplicates().rename(columns={"author_id": "id"})
    
    return result.sort_values("id").reset_index(drop=True)
