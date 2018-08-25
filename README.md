# tiltify
A Tiltify API Library in Go! Work in Progress.


Usage:

```Go
var client = tiltify.Client{}
client.SetAuthKey("AUTH KEY HERE")
client.SetURL("https://tiltify.com/api/v3/")
client.GetCampaignDonations(4814)
 ```

To run tests, you'll need to set an auth key in `campaign_test.go`
