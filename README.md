# Project 5 - 

## A Blockchain Based Voting System
### Why: 
   There is human error as well as corruption in the voting process, and vulnerabilities exist throughout the current voting process from start to end. The transportation and counting of paper ballots leaves the process of counting and recording votes open to serious vulnerabilities. Paper ballots can be directly modified, and whole boxes of votes have been lost or stolen in certain districts. Also, the digitized voting process would be cheaper long term. Cost of paper, mail, and labor for managing polling stations in the current infrastructure will be more than the creation and maintenance of this system. There are similar issues with states that have implemented voting machines, and poor implementations of such machines have already decimated public trust in them.
   
   Digital voting is not feasible without blockchain technology. Having a publicly readable, immutable and verifiable ledger will provide transparency, and thus interest and faith in the voting process. Having a decentralized network of nodes (every polling station can be replaced by a node) confirms that the data cannot be tampered with. In a centralized system, this is not the case. If the one centralized system is compromised, all the voting results are also compromised. 
### How: 
To effectively digitalize the voting process, the system must meet these criteria:

* Integrity: Only eligible voters may vote, and they may only vote once. (Thus at some point there will need to be verification off the chain).
The blockchain will need to have a permissioned ledger so only permissioned nodes (polling stations) have write permissions. Read permissions will be public.
* Transparency: All results should be independently verifiable.
* Privacy: Choices of a voter must be kept private during and after the election. (Prevent voting under duress/buying of votes).

The voter registration process would still have to take place off the grid or off the chain. Once a voter registration agency determines that someone is eligible to vote, they would receive a token or key that would allow them to vote exactly once. This can be distributed digitally, or via snail mail. Either way, since your identity needs to be confirmed at some point during the entire voting process, this is unavoidable. This guarantees the integrity requirement. Also, it is assumed that polling stations will still exist, and voters would have to go to the stations in order to vote. They would just do so digitally. Although it would be ideal to allow the user to vote from the comfort of their own home, that is outside the scope of this project.
### Project Description
Develop the voting system detailed above with the fundamental requirements of Transparency and Privacy. Integrity will be partially guaranteed off chain, and partially guaranteed by having a ledger with write permissions.  

Send: A node will await user input of an ID string from a node allowed (determined by the permissioned ledger). If that ID is confirmed as correct (make an API call to a centralized server that verifies user identity) and the node is permissioned, the node will connect to the network and download the voting ballot. The node will provide a prompt for inputting votes. User will make their decisions and save/confirm. The voting results will be hashed, and the block will be encrypted with the ring signature to maintain anonymity. That will be pushed to the blockchain. 

Receive: The other nodes will need to verify the new block using the key image.
Success Criteria:
* Only specific nodes/roles can write to the blockchain (Permissioned Ledger successfully implemented)
* Blockchain does not fork, or forks are resolved immediately and no data is lost.
* Voting nodes can submit anonymous votes to the blockchain (Ring Signature successfully implemented)
* Votes can be verified on the blockchain (but identities of the submitter remain anonymous)

## Midpoint Review

The original timeline did not allocate time to implement public/private keys or digital signatures.
Ring signature functionality is dependent on having some type of existing digital signature infrastructure in place.



Updated Timeline:
Week 1a: Grant write privileges to blockchain from only certain nodes (Permissioned Ledger) :white_check_mark:

Week 1b: Provide voting user interface and downloadable ballet :white_check_mark:

Week 2a: Implement solution to blockchain forking. (Implement Round Robin Consensus, Delegated Proof of Stake, or something similar) :x:

Week 2b: Implement data transactions and Digital Signatures for validation on blockchain :hourglass_flowing_sand:

CHECKPOINT - 5/1/19

Week 3a: Start work on encrypting voting results using ring signature algorithm

Week 3b: Verify on the blockchain by digital signature

Week 4: Verify integrity of the votes and count the votes
Final Due Date/Demo - 5/16/19 at 2:30 p.m.

Permissioned Blockchain

Voting User Interface and Downloadable Ballot





## Documentation
You will need to launch the initial node on PORT 6688, and then send a GET request to /create the create the first blockchain. 
You will then need to send a GET request to /start to start creating blocks. You can then run and start the rest of the nodes.