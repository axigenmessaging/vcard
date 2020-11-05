"# vcard"

#create a vcard 3.0

### Usage

```
 vc := NewVCardV3()
 prop1 := vc.CreateProperty("name")
 prop1.SetValue(NewText("my name"))

 vc.AddProperty(prop1)

 prop2 := vc.CreateProperty("photo")
 prop2.SetValue(NewText("base64 encoded text"))

 vc.AddPropertyParameter(prop2, "encoding", []string{"base64"})
 vc.AddPropertyParameter(prop2, "type", []string{"jpeg"})

 vc.AddProperty(prop2)

 vcardString := vc.Build()

```

@TO DO: version 4.0, conversions between versions
