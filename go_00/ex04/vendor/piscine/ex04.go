/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ex04.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: miyuu <miyuu@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/28 19:13:46 by miyuu             #+#    #+#             */
/*   Updated: 2024/06/29 21:58:56 by miyuu            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				ft.PrintRune("%c%c%c", a, b, c)
					if a != '7'{
						ft.PrintRune(", ")
				}
			}
		}
	}
	ft.PrintRune()
}
