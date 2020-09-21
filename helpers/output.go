package helpers

import "github.com/b4ysolutions/loki/model"

// Outputs represents the list of return messages
var Outputs = map[string]model.Output{
	"ErrorDBConnection":            model.Output{ID: 1001, Type: "error", Message: "model.Output "},
	"ErrorRetrieveAffectedRows":    model.Output{ID: 1002, Type: "error", Message: "Failed to retrieve affected rows"},
	"ErrorRetrieveLastInsertId":    model.Output{ID: 1003, Type: "error", Message: "Failed to retrieve last insert id"},
	"ErrorExecutePrepareStatement": model.Output{ID: 1004, Type: "error", Message: "Failed to execute prepare statement"},
	"ErrorPreparingDBQuery":        model.Output{ID: 1005, Type: "error", Message: "Failed to preparing database query"},
	"ErrorCreateUser":              model.Output{ID: 1006, Type: "error", Message: "Failed to create the user"},
	"ErrorExecuteQuery":            model.Output{ID: 1007, Type: "error", Message: "Failed to execute database query"},
	"ErrorReadUsersListData":       model.Output{ID: 1008, Type: "error", Message: "Failed to read users list data"},
	"ErrorReadUserData":            model.Output{ID: 1009, Type: "error", Message: "Failed to read user data"},
	"ErrorBindJsonInputData":       model.Output{ID: 1010, Type: "error", Message: "Failed to bind json input data"},
	"SuccessUserRemoved":           model.Output{ID: 2001, Type: "success", Message: "User successfully removed"},
	"SuccessUserUpdated":           model.Output{ID: 2002, Type: "success", Message: "User successfully updated"},
	"SuccessListWithAllUsers":      model.Output{ID: 2003, Type: "success", Message: "List with all users"},
	"SuccessUserFound":             model.Output{ID: 2004, Type: "success", Message: "User found"},
	"SuccessUserCreated":           model.Output{ID: 2005, Type: "success", Message: "User successfully created"},
	"SuccessUserNotRemoved":        model.Output{ID: 2006, Type: "success", Message: "User not removed"},
	"SuccessUserNotUpdated":        model.Output{ID: 2007, Type: "success", Message: "User not updated"},
	"SuccessNoUserFound":           model.Output{ID: 2008, Type: "success", Message: "No user found"},
	"SuccessUserNotFound":          model.Output{ID: 2009, Type: "success", Message: "User not found"},
	"SuccessUserLogin":             model.Output{ID: 2010, Type: "success", Message: "User successfully login"},
	"SuccessUserAlreadyLogged":     model.Output{ID: 2011, Type: "success", Message: "User is already logged"},
	"SuccessUserLogout":            model.Output{ID: 2012, Type: "success", Message: "User successfully logout"},
	"SuccessUserNotLogged":         model.Output{ID: 2013, Type: "success", Message: "User is not logged"},
}

// Output is used to return a success or error json message
func Output(messageCode string, err error) model.Output {
	var output model.Output
	output = Outputs[messageCode]
	output.Details = err
	return output
}

// OutputMessage is used to return only the message
func OutputMessage(messageCode string) string {
	return Outputs[messageCode].Message
}
