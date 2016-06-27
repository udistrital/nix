package utilidades


import(
"log"
"nix/model"
)

func CheckErr(err error, msg string) model.MessageReturn{
    if err != nil {
        log.Println(msg, err)
        return model.GetNewMessage(1,"ERROR: "+msg+" TRACE: "+err.Error())
    }

    return CheckInfo("OK")
}

func CheckInfo(msg string) model.MessageReturn{
	if msg != "OK"{
		log.Println("INFO: "+msg)	
	}
	
    return model.GetNewMessage(0,"INFO: "+msg)
}

