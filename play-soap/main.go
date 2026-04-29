package main

import (
	"encoding/xml"
	"fmt"
)

var xmlExample = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <mitraInfoResponse>
            <return>
                <quota xsi:type="xsd:string">85435906863</quota>
                <Description xsi:type="xsd:string">DEVELMODE BY MUSE</Description>
                <avail_airline xsi:type="SOAP-ENC:Array" SOAP-ENC:arrayType="unnamed_struct_use_soapval[9]">
                    <item>
                        <airline xsi:type="xsd:string">AIR</airline>
                        <label xsi:type="xsd:string">Air Asia</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">AVI</airline>
                        <label xsi:type="xsd:string">Aviastar</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">CIT</airline>
                        <label xsi:type="xsd:string">Citilink</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">EXP</airline>
                        <label xsi:type="xsd:string">Express Air</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">GAR</airline>
                        <label xsi:type="xsd:string">Garuda Indonesia</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">KAL</airline>
                        <label xsi:type="xsd:string">Kalstar</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">LIO</airline>
                        <label xsi:type="xsd:string">Lion Air</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">TRA</airline>
                        <label xsi:type="xsd:string">Transnusa</label>
                    </item>
                    <item>
                        <airline xsi:type="xsd:string">TRI</airline>
                        <label xsi:type="xsd:string">Trigana Air</label>
                    </item>
                </avail_airline>
            </return>
        </mitraInfoResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

var soapErr = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"
  xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
  xmlns:xsd="http://www.w3.org/2001/XMLSchema"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <SOAP-ENV:Fault>
            <faultcode xsi:type="xsd:string">4</faultcode>
            <faultactor xsi:type="xsd:string">credential FAILED</faultactor>
            <faultstring xsi:type="xsd:string">credential FAILED</faultstring>
            <detail xsi:type="xsd:string">credential FAILED |34.101.185.31</detail>
        </SOAP-ENV:Fault>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

var airports = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/">
    <SOAP-ENV:Body>
        <airportCodeResponse>
            <return xsi:type="SOAP-ENC:Array" SOAP-ENC:arrayType="unnamed_struct_use_soapval[418]">
                <item>
                    <value xsi:type="xsd:string">AEG</value>
                    <label xsi:type="xsd:string">Aek Godang (AEG)</label>
                </item>
                <item>
                    <value xsi:type="xsd:string">AGI</value>
                    <label xsi:type="xsd:string">Agimuga (AGI)</label>
                </item>
                <item>
                    <value xsi:type="xsd:string">PQC</value>
                    <label xsi:type="xsd:string">Phu Quoc (PQC)</label>
                </item>
            </return>
        </airportCodeResponse>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

func main() {
	res := &AirportListEnvelope{}
	err := xml.Unmarshal([]byte(airports), res)

	fmt.Println(res.Body, err)
}

type SOAPMerchantBodyResponse struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName xml.Name
	Body    MitraInfoResponse `xml:"mitraInfoResponse"`
}

type MitraInfoResponse struct {
	XMLName xml.Name
	Return  Return `xml:"return"`
}

type Return struct {
	Quota         string       `xml:"quota"`
	Description   string       `xml:"Description`
	AvailAirlines AvailAirline `xml:"avail_airline"`
}

type AvailAirline struct {
	Item []Airline `xml:"item"`
}

type Airline struct {
	Airline string `xml:"airline"`
	Label   string `xml:"label"`
}

type SOAPErrResponse struct {
	XMLName xml.Name
	Body    SOAPErrBOdy `xml:"Body"`
}

type SOAPErrBOdy struct {
	XMLName xml.Name
	Fault   SOAPErrFault `xml:"Fault"`
}

type SOAPErrFault struct {
	XMLName     xml.Name
	FaultCode   string `xml:"faultcode"`
	FaultActor  string `xml:"faultactor"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

type AirportListEnvelope struct {
	XMLName xml.Name
	Body    AirportListBody `xml:"Body"`
}

type AirportListBody struct {
	Response AirportListResponse `xml:"airportCodeResponse"`
}

type AirportListResponse struct {
	Return AirportReturn `xml:"return"`
}

type AirportReturn struct {
	Airports []AirportItem `xml:"item"`
}

type AirportItem struct {
	Value string `xml:"value" json:"code"`
	Label string `xml:"label" json:"label"`
}
