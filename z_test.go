package stringstyles

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("Test started")
	os.Exit(m.Run())
}

func TestSankeCase(t *testing.T) {
	input := "TheQuickBrownFoxJumpsOverTheLazyDog"
	expected := "the_quick_brown_fox_jumps_over_the_lazy_dog"
	output := SnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER_CASE_EMBEDDED_UNDERSCORE"
	expected = "lower_case_embedded_underscore"
	output = SnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "tHeqUicKBrOWnFoXJUmpsoVeRThElAzydOG"
	expected = "t_heq_uic_k_br_o_wn_fo_xj_umpso_ve_r_th_el_azyd_og"
	output = SnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "AAzAzzAAZAazAaz"
	expected = "a_az_azz_aaz_aaz_aaz"
	output = SnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}
func TestScreamingSankeCase(t *testing.T) {
	input := "TheQuickBrownFoxJumpsOverTheLazyDog"
	expected := "THE_QUICK_BROWN_FOX_JUMPS_OVER_THE_LAZY_DOG"
	output := ScreamingSnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "upper_case_embedded_underscore"
	expected = "UPPER_CASE_EMBEDDED_UNDERSCORE"
	output = ScreamingSnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "tHeqUicKBrOWnFoXJUmpsoVeRThElAzydOG"
	expected = "T_HEQ_UIC_K_BR_O_WN_FO_XJ_UMPSO_VE_R_TH_EL_AZYD_OG"
	output = ScreamingSnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "AAzAzzAAZAazAaz"
	expected = "A_AZ_AZZ_AAZ_AAZ_AAZ"
	output = ScreamingSnakeCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}

func TestCamelCase(t *testing.T) {

	input := "the_quick_brown_fox_jumps_over_the_lazy_dog"
	expected := "theQuickBrownFoxJumpsOverTheLazyDog"
	output := CamelCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER_CASE_EMBEDDED_UNDERSCORE"
	expected = "lowerCaseEmbeddedUnderscore"
	output = CamelCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "t_heq_uic_k_br_o_wn_fo_xj_umpso_ve_r_th_el_azyd_og"
	expected = "tHeqUicKBrOWnFoXjUmpsoVeRThElAzydOg"
	output = CamelCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "a_az_azz_aaz_aaz_aaz"
	expected = "aAzAzzAazAazAaz"
	output = CamelCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER_CAsE_eMBEDDED_UNDERSCORe"
	expected = "lowerCAsEEMbeddedUnderscoRe"
	output = CamelCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}

func TestPascalCase(t *testing.T) {

	input := "the_quick_brown_fox_jumps_over_the_lazy_dog"
	expected := "TheQuickBrownFoxJumpsOverTheLazyDog"
	output := PascalCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER_CASE_EMBEDDED_UNDERSCORE"
	expected = "LowerCaseEmbeddedUnderscore"
	output = PascalCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "t_heq_uic_k_br_o_wn_fo_xj_umpso_ve_r_th_el_azyd_og"
	expected = "THeqUicKBrOWnFoXjUmpsoVeRThElAzydOg"
	output = PascalCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "a_az_azz_aaz_aaz_aaz"
	expected = "AAzAzzAazAazAaz"
	output = PascalCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER_CAsE_eMBEDDED_UNDERSCORe"
	expected = "LowerCAsEEMbeddedUnderscoRe"
	output = PascalCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}

func TestKebabCase(t *testing.T) {
	input := "TheQuickBrownFoxJumpsOverTheLazyDog"
	expected := "the-quick-brown-fox-jumps-over-the-lazy-dog"
	output := KebabCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "LOWER-CASE-EMBEDDED-UNDERSCORE"
	expected = "lower-case-embedded-underscore"
	output = KebabCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "tHeqUicKBrOWnFoXJUmpsoVeRThElAzydOG"
	expected = "t-heq-uic-k-br-o-wn-fo-xj-umpso-ve-r-th-el-azyd-og"
	output = KebabCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}

	input = "AAzAzzAAZAazAaz"
	expected = "a-az-azz-aaz-aaz-aaz"
	output = KebabCase(input)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}

func TestStudlyCaps(t *testing.T) {
	input := "TheQuickBrownFoxJumpsOverTheLazyDog"
	expected := "ThEquICkbROwNFoXJumPSOVErTHElazyDOg"
	output := StudlyCaps(input, 31283726224234348)
	if output != expected {
		t.Errorf("conver %s failed: want %s got %s", input, expected, output)
	}
}
