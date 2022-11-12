```
 /************************** WERTETABELLE *******************************/

    /* speedOnSurface Daten aus der Wertetabelle */
    private double[] speedOnSurfaceTable = { 0, 3, 5, 7.5, 8, 11, 14, 18,
            20, 22 };
    /* fuelPerHour Daten aus der Wertetabelle */
    private double[] fuelPerHourTable = { 80, 90, 105, 140, 131, 162, 197,
            251, 280, 310 };

    /*
     * Die Wertetabelle wird in dieser Methode als lineares Spline approximiert.
     */
    public double getFuelPerHourLinear(double speedOnSurface) {
        //check definitionsmenge
    	if(speedOnSurface < 0) {
    		throw new IllegalArgumentException("speedOnSurface ist negativ");
    	}
    	if(speedOnSurface > speedOnSurfaceTable[speedOnSurfaceTable.length-1]) {
    		throw new IllegalArgumentException("speedOnSurface ist höher als der höchste Referenzwert");
    	}
    	
    	//get the arg, val of the surrounding elements
    	//for aproximation
    	//if speedOnSurface is in the table (= arg) then return the value
    	int i = 0;
    	for(;i < speedOnSurfaceTable.length; i++) {
    		if(speedOnSurfaceTable[i] == speedOnSurface) {
    			return fuelPerHourTable[i];
    		} else if(speedOnSurfaceTable[i] < speedOnSurface) {
    			//index is at the nearest element lower than speedOnSurface
    			break;
    		}
    	}
    	double fun_argx1 = speedOnSurfaceTable[i]; //lower fun x
    	double fun_argx2 = speedOnSurfaceTable[i+1]; //higher fun x
    	
    	double fun_valy1 = fuelPerHourTable[i]; //lower fun y
    	double fun_valy2 = fuelPerHourTable[i+1]; //higher fun y
    	
    	// f: y = a * x + b
    	// fun_valy1 = a * fun_argx1 + b 
    	//--> y - ax = b
    	// fun_valy2 = a * fun_argx2 + b
    	//--> fun_valy2 = a * fun_argx2 + fun_valy1 - a * fun_argx1
    	// fun_valy2 - fun_valy1 = a * fun_argx2 - a * fun_argx1
    	// fun_valy2 - fun_valy1 = a * (fun_argx2 - fun_argx1)
    	// (fun_valy2 - fun_valy1)/(fun_argx2 - fun_argx1) = a
    	// y - ax = b --> y - ((fun_valy2 - fun_valy1)/(fun_argx2 - fun_argx1))*x = b
    	
    	double a = (fun_valy2 - fun_valy1)/(fun_argx2 - fun_argx1);
    	double b = fun_valy1 - (a*fun_argx1);
    	// f: x |-> a*x+a
    	return a*speedOnSurface+b;
    }

```
