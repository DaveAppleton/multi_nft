// ------------ THE ORIGINAL ---------------------

let showMintButton = isEmptyObject(currentOffer) && isEmptyObject(originalOffer) && resProject.ownerDetails.address === userData.address && resProject.approved;
let showBuyButton = !isEmptyObject(currentOffer) && 
                    currentOffer.creator !== userData.address && 
                    currentOffer.balanceOf > 0;
let showBuyToNonPatron = !isEmptyObject(currentOffer) && currentOffer.creator === originalOffer.creator && currentOffer.pos === 0 && originalOffer.pos === 0 && isPatron;
let showResaleButton = isEmptyObject(currentOffer) && !isEmptyObject(originalOffer) && resProjectArt.ownerdetails.address === userData.address;

//---- Made a LOT clearer

// 1.0 explain all the conditions

let currentOfferDoesNotExist = isEmptyObject(currentOffer)
let currentOfferExists = !currentOfferDoesNotExist
let originalOfferDoesNotExist = isEmptyObject(originalOffer)
let viewerIsOwner = resProject.ownerDetails.address === userData.address;
let viewerIsNotOwner = !viewerIsOwner
let someLeftToSell = currentOfferExists && currentOffer.balanceOf > 0     // offer must exist to avoid NULL errors
let isMinting = currentOfferExists && currentOffer.pos === 0              // offer must exist to avoid NULL errors

// 2.0 now you use easy to use terms to create boolean logic

let showMintButton = currentOfferDoesNotExist &&                           // no current offer ????
                     originalOfferDoesNotExist  &&                         // no original offer <-- isn't this enough ?
                     viewerIsOwner &&                                      // you are the creator
                     resProject.approved;                                  // project is approved

// ------- this is not well thought out

let showBuyButton = viewerIsNotOwner   &&                         // you are not the person who put it on sale
                    someLeftToSell;                               // there are some left to sell

// this covers ALL resale cases

let showBuyToNonPatron = someLeftToSell && 
                        currentOffer.creator === originalOffer.creator &&   // the offer we are looking at exists
                        originalOffer.pos === 0 &&                          // why?? the original offer MUST be zer0 (or it isn't the first)
                        (
                            (isPatron && currentOffer.pos === 0)||  // ONLY PATRONS CAN BUY POSITION ZERO (first sale)
                            (currentOffer.pos > 0)                  // anybody can buy resale offers
                        ) &&
                        viewerIsNotOwner;                           // No point asking owner to buy
                                        


let showResaleButton = isEmptyObject(currentOffer) &&                       // the offer we are looking does not exist ?
                    !isEmptyObject(originalOffer) &&                        // there WAS an original offer
                    resProjectArt.ownerdetails.address === userData.address;// you are the owner

