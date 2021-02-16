from decimal import *
import time


def trial_division(n):
	t = time.process_time_ns()

	p_list = []
	p = Decimal(2)

	it_count = 1
	while (p*p <= n):
		while (n%p == 0):
			p_list.append(p)
			n /= p

		p += 1
		
		it_count += 1
	
	if (n != 1):
		p_list.append(n)

	if (len(p_list) > 2):
	    print("AAAAAA")
	
	return int(p_list[0]), int(p_list[1]), it_count, time.process_time_ns()-t




def trial_division_improved(n):
	t = time.process_time_ns()

	p_list = []

	it_count = 0

	while (n % 2 == 0):
	    p_list.append(2)
	    n /= 2

	    it_count += 1

	p = Decimal(3)

	while (p*p <= n):
		while (n%p == 0):
			p_list.append(p)
			n /= p

		p += 2
		
		it_count += 1
	
	if (n != 1):
		p_list.append(n)

	if (len(p_list) > 2):
	    print("AAAAAA")
	
	return int(p_list[0]), int(p_list[1]), it_count, time.process_time_ns()-t