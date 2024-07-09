/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex05.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: miyuu <miyuu@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/28 19:13:46 by miyuu             #+#    #+#             */
/*   Updated: 2024/06/29 22:02:24 by miyuu            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '8'; b++ {
			for c := a; c <= '9'; c++ {
				for d := b + 1; d <= '9'; d++ {
						ft.PrintRune(a)
						ft.PrintRune(b)
						ft.PrintRune(' ')
						ft.PrintRune(c)
						ft.PrintRune(d)
						if !(a == '9' && b == '8' ){
							ft.PrintRune(',')
							ft.PrintRune(' ')
					}
				}
			}
		}
	}
	ft.PrintRune('\n')
}
