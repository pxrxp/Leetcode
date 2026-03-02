import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.functions.length, col

object Solution {
  def invalidTweets(tweets: DataFrame): DataFrame = {
    tweets
      .filter(length(col("content")) > 15)
      .select("tweet_id")
  }
}
