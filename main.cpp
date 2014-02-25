/*
 * main.cpp
 *
 *  Created on: Aug 7, 2009
 *      Author: Alex Swehla
 */

#include <Cmath>
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;


class QMatrix {

	private:
		//vector used as Dynamic Array
		vector<int> gMatrix;

		//Store value of N in class for computation
		int m_nNumQ;
	public:
		//Constructor
		QMatrix (long int N=8) {
			gMatrix.resize(N,-1);
			m_nNumQ= (int) N;
		}

		//Calculate Solution to N Queens Problem Given N
		int GetSolution(int nPos) {
			int nErr=1;
			bool bFail=false;
			gMatrix[nPos]= -1;

			do {

				//increment current column position
				gMatrix[nPos]++;

					//test for row 1 if row 1 move to next row
					if (nPos==0){	nErr=GetSolution(nPos+1);}
					else {

						bFail=false;
						//for all other rows test piece placement if safe move to next row
						for (int i=1;i<=nPos;i++) {

									//fail entire tree from current node if position not safe
									if(gMatrix[nPos]==gMatrix[nPos-i] || fabs(gMatrix[nPos]-gMatrix[nPos-i])==fabs(nPos-(nPos-i))){
										bFail=true;
										break;
									}
								}

								if(bFail==false){
									//test for end row, if end row return successful
									if(nPos<(m_nNumQ-1)){
										//move to next row
										nErr=GetSolution(nPos+1);
									}
									else {
										return 0;
									}
								}
					}
			}while (gMatrix[nPos]< (m_nNumQ-1) && nErr==1);

			return nErr;
		}

		//Output Solution to file
		int PrintSolution(){

			//Open File for Writing
			ofstream outf("output.txt");

			//Fail if File could not be opened
			if(!outf){
				return 1;
			}

			int nX, nY;

			//Print solution to file
			outf<<"Solution for N="<<m_nNumQ<<endl;
			for (nY=0;nY<m_nNumQ;nY++){
				for (nX=0;nX<m_nNumQ;nX++){
					if (gMatrix[nY]==nX){
						outf<<"|Q";
					}
					else {
						outf<<"|_";
					}
				}
				outf<<"|"<<endl;
			}
			return 0;
		}
};

int main(int argc, char** argv){

	// If the user didn't provide a value for N command line argument,
	// print an error and exit.
	    if (argc <= 1)
	    {
	    	cout<<"Please Specify an Integer Value for 'N'"<<endl;
	        cout<<"Usage: "<<argv[0]<<" <Value of N> "<<endl;
	        exit(1);
		}

	long int N=strtol(argv[1], NULL ,10);
	int nErr;

	QMatrix nMatrix(N);

	nMatrix.GetSolution(0);

	nErr=nMatrix.PrintSolution();

	if(nErr==1){
		cout<<"ERROR: Could Not Write to File"<<endl;
	}

	return 0;
}

