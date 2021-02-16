#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import click

from trial_division import *
from fermat import *
from utils import *

from decimal import *
from tabulate import tabulate


def factorize(n, method):

	n = Decimal(n)

	headers = ["Method", "Result", "Exec Time (ns)", "nIteration", "Verification"]
	results = []

	if method == "all":
		factorize(n, "trial")
		factorize(n, "fermat")

	if method == "trial":

		p, q, itc, t = trial_division(n)
		results.append(["Trial", str((p, q)), t, itc, check(p, q, n)])

		p, q, itc, t = trial_division_improved(n)
		results.append(["Trial v2", str((p, q)), t, itc, check(p, q, n)])

	if method == "fermat":
		p, q, itc, t = fermat(n)
		results.append(["Fermat", str((p, q)), t, itc, check(p, q, n)])

	if method != "all":
		print(tabulate(results, headers, tablefmt="fancy_grid"))



# program arguments handling
def handleArgs(p, q, n, method):
	print("")
	if p and q and n:
		print("Please provide ether [--n] or [--p and --q]")
		return 1

	if p and q:
		print(f'N = p * q = {p} * {q} = {p*q}')
		print("Factorizing N =", p*q)
		factorize(p*q, method)
		print()
		return 0

	if n:
		print("Factorizing N =", n)
		factorize(n, method)
		print()
		return 0

	else:
		print("Please provide ether [--n] or [--p and --q]")


@click.command()
@click.option("--p", type=click.INT, help="Specify a p like N=p*q (--q is required)")
@click.option("--q", type=click.INT, help="Specify a q like N=p*q (--p is required)")
@click.option("--n", type=click.INT, help="Specify the N number to factor (must be a composite number)")

@click.option("--method", default="all", type=click.Choice(['all', 'trial', 'fermat'], case_sensitive=False), help="Specify the N number to factor (must be a composite number)")

def main(p, q, n, method):
    """ Main program """

    handleArgs(p, q, n, method)

    return 0

if __name__ == "__main__":
    main()