package v1alpha1

import (
	"kubevault.dev/apimachinery/apis/kubevault/v1alpha2"

	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1alpha1_MySQLSpec_To_v1alpha2_MySQLSpec(in *MySQLSpec, out *v1alpha2.MySQLSpec, s conversion.Scope) error {
	out.Address = in.Address
	out.Database = in.Database
	out.Table = in.Table
	out.CredentialSecretRef = in.UserCredentialSecret
	out.TLSSecretRef = in.TLSCASecret
	// WARNING: in.UserCredentialSecret requires manual conversion: does not exist in peer-type
	// WARNING: in.TLSCASecret requires manual conversion: does not exist in peer-type
	out.MaxParallel = in.MaxParallel
	return nil
}

func Convert_v1alpha2_MySQLSpec_To_v1alpha1_MySQLSpec(in *v1alpha2.MySQLSpec, out *MySQLSpec, s conversion.Scope) error {
	out.Address = in.Address
	out.Database = in.Database
	out.Table = in.Table
	out.UserCredentialSecret = in.CredentialSecretRef
	out.TLSCASecret = in.TLSSecretRef
	// WARNING: in.CredentialSecretRef requires manual conversion: does not exist in peer-type
	// WARNING: in.TLSSecretRef requires manual conversion: does not exist in peer-type
	out.MaxParallel = in.MaxParallel
	// WARNING: in.DatabaseRef requires manual conversion: does not exist in peer-type
	// WARNING: in.PlaintextCredentialTransmission requires manual conversion: does not exist in peer-type
	// WARNING: in.MaxIdleConnection requires manual conversion: does not exist in peer-type
	// WARNING: in.MaxConnectionLifetime requires manual conversion: does not exist in peer-type
	// WARNING: in.HAEnabled requires manual conversion: does not exist in peer-type
	// WARNING: in.LockTable requires manual conversion: does not exist in peer-type
	return nil
}

func Convert_v1alpha1_PostgreSQLSpec_To_v1alpha2_PostgreSQLSpec(in *PostgreSQLSpec, out *v1alpha2.PostgreSQLSpec, s conversion.Scope) error {
	out.CredentialSecretRef = in.ConnectionURLSecret
	// WARNING: in.ConnectionURLSecret requires manual conversion: does not exist in peer-type
	out.Table = in.Table
	out.MaxParallel = in.MaxParallel
	return nil
}

func Convert_v1alpha2_PostgreSQLSpec_To_v1alpha1_PostgreSQLSpec(in *v1alpha2.PostgreSQLSpec, out *PostgreSQLSpec, s conversion.Scope) error {
	out.ConnectionURLSecret = in.CredentialSecretRef
	// WARNING: in.CredentialSecretRef requires manual conversion: does not exist in peer-type
	out.Table = in.Table
	out.MaxParallel = in.MaxParallel
	// WARNING: in.MaxIdleConnection requires manual conversion: does not exist in peer-type
	// WARNING: in.HAEnabled requires manual conversion: does not exist in peer-type
	// WARNING: in.HaTable requires manual conversion: does not exist in peer-type
	return nil
}
