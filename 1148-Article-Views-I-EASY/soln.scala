import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.functions.col

object Solution {
  def selfViews(views: DataFrame): DataFrame = {
    views
      .filter(col("author_id") === col("viewer_id"))
      .select(col("author_id").alias("id"))
      .distinct()
      .orderBy("id")
  }
}
