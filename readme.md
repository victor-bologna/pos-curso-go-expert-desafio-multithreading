# Multithreading Request CEP API:

Simple program do fetch Brazil Address information from two sources by CEP at the same time to see who has the fastest response via golang's channel and go routine to. 

If both requests fail to return the data within 1 second then the program returns Timeout.