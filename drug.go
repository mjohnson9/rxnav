package rxnav

import (
	"encoding/xml"
	"errors"
)

// Drug represents a single drug, identified by its RxCUI

// Drug represents a single drug, identified by its RxCUI

// Drug represents a single drug, identified by its RxCUI

// Drug represents a single drug, identified by its RxCUI
type Drug struct {
	RxCUI int
	Name  string
}

var ErrNoMatches = errors.New("rxnav: no matches were found for the given drug ID")

// Available ID lookup types
const (
	IDGoldStandard                        = "AMPID"
	IDAbbreviatedNewAnimalDrugApplication = "ANADA"
	IDAbbreviatedNewDrugApplication       = "ANDA"
	IDATC                                 = "ATC"
	IDBiologicsLicenseApplication         = "BLA"
	IDVaccine                             = "CVX"
	IDDrugBank                            = "Drugbank"
	IDGCNSeqNo                            = "GCN_SEQNO"
	IDGenericFormulaCode                  = "GFC"
	IDHCPCS                               = "HCPCS"
	IDHICSeqNo                            = "HIC_SEQN"
	IDMeSH                                = "MESH"
	IDMMSLCode                            = "MMSL_CODE"
	IDNewAnimalDrugApplication            = "NADA"
	IDNewDrugApplication                  = "NDA"
	IDNationalDrugCode                    = "NDC"
	IDSNOMED                              = "SNOMEDCT"
	IDSPLSetID                            = "SPL_SET_ID"
	IDUMLSCUI                             = "UMLSCUI"
	IDUNIICode                            = "UNII_CODE"
	IDUSP                                 = "USP"
	IDVAUniqueID                          = "VUID"
)

type drugByIDResponse struct {
	XMLName xml.Name `xml:"rxnormdata"`
}

func GetDrugByID(idType string, id string) (*Drug, error) {

}
