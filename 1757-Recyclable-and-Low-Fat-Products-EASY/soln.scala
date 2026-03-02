import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.functions.col

object Solution {
  def findProducts(products: DataFrame): DataFrame = {
    products
      .filter(col("low_fats") === "Y" && col("recyclable") === "Y")
      .select("product_id")
  }

}
