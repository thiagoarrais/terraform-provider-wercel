package wercel

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/thiagoarrais/terraform-provider-wercel/sdk"
)

func errorFromSDKErr(err error) error {
	if sdkErr, ok := err.(sdk.GenericOpenAPIError); ok {
		if errResp, ok := sdkErr.Model().(sdk.ErrorResponse); ok {
			if errResp.HasError() && errResp.Error.HasMessage() {
				return fmt.Errorf(errResp.Error.GetMessage())
			}
		}
	}
	return err
}

func diagFromSDKErr(err error) diag.Diagnostics {
	return diag.FromErr(errorFromSDKErr(err))
}
