import org.apache.spark.sql.DataFrame
import org.apache.spark.sql.functions.col

object Solution {
  def findCustomerReferee(customer: DataFrame): DataFrame = {
    customer
      .filter(col("referee_id").isNull || col("referee_id") =!= 2)
      .select("name")
  }
}
