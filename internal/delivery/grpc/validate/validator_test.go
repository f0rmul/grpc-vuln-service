package validate

import (
	"testing"

	"github.com/f0rmul/vuln-service/pkg/netvuln_v1"
	"github.com/stretchr/testify/require"
)

func TestValidateWithNoErrors(t *testing.T) {
	request := &netvuln_v1.CheckVulnRequest{
		Targets:  []string{"nmap.org", "128.1.1.1", "myhost.com"},
		TcpPorts: []int32{65535, 22, 111, 1, 33},
	}

	err := ValidateProtoRequest(request)

	require.NoError(t, err)
}

func TestValidateWithErrors(t *testing.T) {
	testCases := []struct {
		testName    string
		input       *netvuln_v1.CheckVulnRequest
		expectedErr error
	}{
		{
			testName: "Invalid target",
			input: &netvuln_v1.CheckVulnRequest{
				Targets:  []string{"http://nmap.org"}, //FIX: CANT CHECK WHATEVER HOST IS IPv4
				TcpPorts: []int32{33, 2, 67, 222},
			},
			expectedErr: ErrInvalidTarget,
		},
		{
			testName: "Invalid port",
			input: &netvuln_v1.CheckVulnRequest{
				Targets:  []string{"nmap.org"},
				TcpPorts: []int32{22, 80, 1234, -1, 65535},
			},
			expectedErr: ErrInvalidPort,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			err := ValidateProtoRequest(testCase.input)
			require.Error(t, err)
			require.EqualError(t, testCase.expectedErr, err.Error())
		})
	}
}
