#include "lowlevel.h"
#include <IOKit/pwr_mgt/IOPMLib.h>
#include <IOKit/IOKitLib.h>

void wakeDisplay(void)
{
    static IOPMAssertionID assertionID;
    IOPMAssertionDeclareUserActivity(CFSTR("MacWakeDisplay"), kIOPMUserActiveLocal, &assertionID);
}
