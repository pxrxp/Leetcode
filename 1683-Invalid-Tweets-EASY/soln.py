import pandas as pd

def invalid_tweets(tweets: pd.DataFrame) -> pd.DataFrame:
    mask = tweets["content"].str.len() > 15
    cols = ["tweet_id"]
    
    result = tweets.loc[mask, cols]
    return result.reset_index(drop=True)
