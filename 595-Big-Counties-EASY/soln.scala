import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.functions.col

object Solution {
  def bigCountries(world: DataFrame): DataFrame = {
    world
      .filter(col("area") >= 3000000 || col("population") >= 25000000)
      .select("name", "population", "area")
  }
}
