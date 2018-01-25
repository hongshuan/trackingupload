package main

type DeliveryRequest struct {
    DeliveryRequestId string `xml:"delivery-request-id"`
    JobId             string `xml:"job-id"`
    Status            string `xml:"status"`
    Category          string `xml:"category"`

    type DeliverySpec struct {
        type Detination struct {
            type Recipient struct {
                ContactName       string `xml:"contact-name"`
                AddressLine1      string `xml:"address-line-1"`
                City              string `xml:"city"`
                ProvState         string `xml:"prov-state"`
                PostalZipCode     string `xml:"postal-zip-code"`
                CountryCode       string `xml:"country-code"`
                ClientVoiceNumber string `xml:"client-voice-number"`
            }
        }

        LabelCount string `xml:"label-count"`
        ProductId  string `xml:"product-id"`

        type Sender struct {
            ContactName       string `xml:"contact-name"`
            Company           string `xml:"company"`
            AddressLine1      string `xml:"address-line-1"`
            AddressLine2      string `xml:"address-line-2"`
            City              string `xml:"city"`
            ProvState         string `xml:"prov-state"`
            PostalZipCode     string `xml:"postal-zip-code"`
            CountryCode       string `xml:"country-code"`
            ClientVoiceNumber string `xml:"client-voice-number"`
        }

        type Option struct {
            Code string `xml:"code", attr`
            Amount string `xml:"amount"`
        }

        type ItemSpecification struct {
            Weight            string `xml:"physical-characteristics>weight"`
            ItemType          string `xml:"item-type"`
            CustomerLabelType string `xml:"customer-label-type"`
            Quantity          string `xml:"quantity"`
        }

        type Notification struct {
            EmailSubject string `xml:"email-subject"`
            type MailedByNotifEmail struct {
                Email       string `xml:"email"`
                OnShipment  string `xml:"on-shipment"`
                OnException string `xml:"on-exception"`
                OnDelivery  string `xml:"on-delivery"`
            }
            type MoboNotifEmail struct {
                Email       string `xml:"email"`
                OnShipment  string `xml:"on-shipment"`
                OnException string `xml:"on-exception"`
                OnDelivery  string `xml:"on-delivery"`
            }
            type ClientNotifEmail struct {
                OnShipment  string `xml:"on-shipment"`
                OnException string `xml:"on-exception"`
                OnDelivery  string `xml:"on-delivery"`
            }
            type ClientNotifEmail2 struct {
                OnShipment  string `xml:"on-shipment"`
                OnException string `xml:"on-exception"`
                OnDelivery  string `xml:"on-delivery"`
            }
        }

        type Reference struct {
            ItemId                 string `xml:"item-id"`
            CustomerSuppliedItemId string `xml:"customer-supplied-item-id"`
            CustomerRef1           string `xml:"customer-ref1"`
            Barcode                string `xml:"barcode"`
        }

        EarlyDeposit string `xml:"deposit-specification>early-deposit"`
    }

    ContinuousInboundFreight string `xml:"continuous-inbound-freight"`

    type SettlementDetails struct {
        struct Payment struct {
            Method string `xml:"method", attr`
        }
        MailedByCustomer         string `xml:"mailed-by-customer"`
        MailedOnBehalfOfCustomer string `xml:"mailed-on-behalf-of-customer"`
        PaidByCustomer           string `xml:"paid-by-customer"`
        ContractId               string `xml:"contract-id"`
        MailingDate              string `xml:"mailing-date"`
        Interliner               string `xml:"interliner"`
        MeterAtFullPrice         string `xml:"meter-at-full-price"`
    }

    type PricingResult struct {
        BaseAmount   string `xml:"base-amount"`
        GstAmount    string `xml:"gst-amount"`
        PstAmount    string `xml:"pst-amount"`
        HstAmount    string `xml:"hst-amount"`
        PreTaxAmount string `xml:"pre-tax-amount"`
        DueAmount    string `xml:"due-amount"`
        RateCode     string `xml:"rate-code"`
        RateTaxCode  string `xml:"rate-tax-code"`

        <options>
          <option code="COV">
            <amount>0.00</amount>
          </option>
          <option code="DC">
            <amount>0</amount>
          </option>
        </options>
        <adjustments>
          <option code="FUELSC">
            <amount>0.90</amount>
          </option>
        </adjustments>

        type service-standard struct {
            AmDelivery         string `xml:"am-delivery"`
            GuaranteedDelivery string `xml:"guaranteed-delivery"`
            MinDays            string `xml:"min-days"`
            MaxDays            string `xml:"max-days"`
        }
    }

    type InductionSpec struct {
        OutletId         string `xml:"outlet-id"`
        OutletPostalCode string `xml:"outlet-postal-code"`
    }

    type HandlingPreferences struct {
        ReturnLabelRequested          string `xml:"return-label-requested"`
        LabelInstructionIncluded      string `xml:"label-instruction-included"`
        InsuredValueIncluded          string `xml:"insured-value-included"`
        PostageRateIncluded           string `xml:"postage-rate-included"`
        CostCenterRequired            string `xml:"cost-center-required"`
        AdditionalAddressRequired     string `xml:"additional-address-required"`
        IncludeShippingCost           string `xml:"include-shipping-cost"`
        OrderIdRef1Required           string `xml:"order-id-ref1-required"`
        PackagingInstructionsIncluded string `xml:"packaging-instructions-included"`
    }

    type ReconcileDetails struct {
        PoNumber      string `xml:"po-number"`
        JobName       string `xml:"job-name"`
        ReconcileDate string `xml:"reconcile-date"`
    }
}
