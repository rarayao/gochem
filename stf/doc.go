/*
 * doc.go, part of gochem.
 *
 * Copyright 2021 Raul Mera <rauldotmeraatusachdotcl>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation; either version 2.1 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General
 * Public License along with this program.  If not, see
 * <http://www.gnu.org/licenses/>.
 *
 *
 * /

//package stf implements the simple trajectory format, an internal trajectory format for goChem.
//stf aimst to produce reasonably small files and to be very easy to read and write, so readers/writers
//can be easily implemented in other programing languages / for other libraries or programs.

/******************** Format Specification   ***************************************************

An STF file might have 3 file extension, which indicate
how they are compressed:
	stf: Compressed with deflate, at any level of compression.
	stz: Compressed with gzip (no idea why this is different than stf, but gunzip won't open my deflate-compressed files
	stl: Compressed with lzw

This spec does not establish what to do in case of a filename not ending with f, z or l, but the reference implementation
attempts to open anyway, and assumes it to be compressed with deflate. I repeat, that behavior is _not_ part of the spec, and might change.

A STF file may only contain ASCII symbols.

A STF file has a "header" starting in the first line, and ending with a line that contains only "**".
Each line of the header must be a pair key=value

After that, the file has one line per atom, per frame. Each line 3 numbers (x y z coordinates, in A). The precision is not specified.
each frame ends with a line containing only "*".

***************************************************************************************************/

package stf
