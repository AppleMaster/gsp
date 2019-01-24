resource "null_resource" "create_lambda_zip_file" {
  provisioner "local-exec" {
    command = "cd ${path.module}; rm -f *.zip; zip -q -r lambda_log_forwarder.zip lib index.js"
  }
}

resource "aws_lambda_function" "lambda_log_forwarder" {
  count         = "${var.enabled == 0 ? 0 : 1}"
  depends_on    = ["null_resource.create_lambda_zip_file"]
  filename      = "${path.module}/lambda_log_forwarder.zip"
  function_name = "${var.cluster_name}_log_forwarder"
  role          = "${aws_iam_role.lambda_log_forwarder.arn}"
  handler       = "index.handler"
  runtime       = "nodejs6.10"
  timeout       = "10"
  memory_size   = "128"
  description   = "A function to forward logs from AWS to a Splunk HEC"

  environment {
    variables = {
      SPLUNK_HEC_TOKEN = "${var.splunk_hec_token}"
      SPLUNK_HEC_URL   = "${var.splunk_hec_url}"
    }
  }
}

resource "aws_cloudwatch_log_group" "lambda_log_forwarder" {
  count             = "${var.enabled == 0 ? 0 : 1}"
  name              = "/aws/lambda/${aws_lambda_function.lambda_log_forwarder.function_name}"
  retention_in_days = 7
}

resource "aws_lambda_permission" "cloudwatch_splunk_logs" {
  count         = "${var.enabled == 0 ? 0 : 1}"
  statement_id  = "${var.cluster_name}_cloudwatch_splunk_logs"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.lambda_log_forwarder.arn}"
  principal     = "logs.eu-west-2.amazonaws.com"
  source_arn    = "${var.cloudwatch_log_group_arn}"
}

resource "aws_cloudwatch_log_subscription_filter" "cloudwatch_splunk_logs" {
  count           = "${var.enabled == 0 ? 0 : 1}"
  depends_on      = ["aws_lambda_permission.cloudwatch_splunk_logs"]
  name            = "${var.cluster_name}_cloudwatch_splunk_logs_subscription_filter"
  destination_arn = "${aws_lambda_function.lambda_log_forwarder.arn}"
  filter_pattern  = ""
  log_group_name  = "${var.cloudwatch_log_group_name}"
}
