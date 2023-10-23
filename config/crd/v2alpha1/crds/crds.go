/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcY_o\xe4\xb8\r\x7fϧ \xb6\x0fyY;\xb7ע(\xe6mw\xd2\x02Ao\xd3\xc1e\x91w٢=\xbaȒ*ɓ\xa6E\xbf{A\xc9\xf2\xf8\x8f&N\xf6z緑(\xf2'\x92\xfa\x91\xd2\x14EqŌxD\xeb\x84V;`F\xe0\xbf<*\xfa\xe5ʧ\xbf\xb8R\xe8\x9bӧ\xab'\xa1\xf8\x0e\xf6\xbd\xf3\xba\xfb\x19\x9d\xeem\x8d\xb7\xd8\b%\xbc\xd0\xea\xaaC\xcf8\xf3lw\x05\xc0\x94ҞѰ\xa3\x9f\x00\xb5V\xdej)\xd1\x16-\xaa\U000a9bf0\xea\x85\xe4h\x83\xf2d\xfa\xf4C\xf9\xe9\xc7\xf2\x87+\x00\xc5:\xdc\x01\xe9\xe3\xfaYI\u0378+O(\xd1\xeaR\xe8+g\xb0&ŭս\xd9\xc1y\".\x1c\x8cF\xc0\xb7̳\xdbAG\x18\x96\xc2\xf9\xbf\xaf\xa6~\x12·i#{\xcb\xe4\xc2v\x98qB\xb5\xbddv>w\x05\xe0jmp\a\xf7dڰ\x1ail\xd8S\x80R\x00\xe3<x\x89Ƀ\x15ʣ\xddk\xd9w\xc9;\x05pt\xb5\x15\xc6\a/La\x81\xf3\xcc\xf7\x0e\\_\x1f\x819\xb8\xc7\xe7\x9b;u\xb0\xba\xb5\xe8\",\x80_\x9cV\a\xe6\x8f;(\xa3xi\x8e\xcc\xe10\x1b]\xf9\x10&\x86!\xffBx\x9d\xb7B\xb59\x04\xdfD\x87\xc0{\x1bBH\xfb\xae\x11\xfcQ\xb89\xb4g\xe6\b\x9e\xf5\xc8/\x02\t\xf3\xa4\xcey֙%\xa2\xc9\xd2\b\x893\x8f9@{\xdd\x19\x89\x1e9T/\x1e\xd36\x1am;\xe6w \x94\xff\xf3\x9f.\xfbbpV\x19\x96\xdej5w\xcc\x17\x1a\x85\xc9pDBQj\xd1f\xbd\xa3=\x93\xbf\x06\x88'\x05_&\xeb#\x92\xa8w:\xbe\t\x85R\x0et\x03\xfe\x88\xf0\x85\xd5O\xbd\x81\a\xaf-k\x11~\xd2u\f\xdf\xf3\x11-\x06\x89*JP\xf6\x82\xa0\xd8i\x9b\r\x9d\xc1\xba\x8c\xb2\x83\xb2\xa4k\x11\xbf\xb9\xa1\xff{n\xd5\x16Y6\xb7\x12ՔABh\x95O\xb0\xcf-\xbe)\xb9\xa6NT\x9a\xe3\xc4c3L\u0081\xb1\xbaF\xe7^IxR0Cq\x7f\x1eX\xb9&J\x9c~d\xd2\x1c٧H2\xf5\x11;\xb6\x1bVh\x83\xea\xf3\xe1\xee\xf1\x8f\x0f\xb3a  \x06\xad\x17\x89\xeb\xe27\xa1\xf1\xc9(̷{M\n\xa3\x14p\xe2ota\xaf\x03c!\x1f0D\x97\b\a\x16\x8dE\x87\xcaOÜ>\xdd\x00S\xa0\xab_\xb0\xf6%<\xa0%5\xe0\x8e\xba\x97\x9ch\xff\x84փ\xc5Z\xb7J\xfc{\xd4\xed\xc0\xeb`T2\x8f\x03\xf1\x9e\xbf\xc0\x90\x8aI81\xd9\xe3G`\x8aC\xc7^\xc0\"Y\x81^M\xf4\x05\x11W\xc2Wm\x11\x84j\xf4\x0e\x8e\xde\x1b\xb7\xbb\xb9i\x85O\xe5\xab\xd6]\xd7+\xe1_nB%\x12U\xef\xb5u7\x1cO(o\x9ch\vf\xeb\xa3\xf0X\xfb\xde\xe2\r3\xa2\b\xd0U(ae\xc7\xff`\x87\x82\xe7\xaegXW\x11\x8d_\xa8<\xafD\x80\xca\x0f\xa5\x13\x1b\x96\xc6]\x9c\x1dMC䝟\xff\xfa\xf0\r\x92\xe9\x10\x8c\xa5\xf7\x83\xdf\xcf\v\xdd9\x04\xe40\xa1\x1a\xb41\x88\x8d\xd5]Љ\x8a\x1b-\x94\x0f?j)P-\xdd\xef\xfa\xaa\x13\x9e\xe2\xfe\xcf\x1e\x9d\xa7X\x95\xb0\x0f5\x1d*\x84\xde\xd0A\xe2%\xdc)س\x0e\xe5\x9e9\xfc\xcd\x03@\x9ev\x059\xf6m!\x98\xb6#K\xe1\xe8\xb5\xc9D\xea'.\xc4k\xca\x02\x0f\x06k\n\x1dy\x8f\x96\x89F\f<\xdbh\vl&[\xceT\xe6\x8f,}Y\xae]\n-0}ɭI\xc0Ԅ\xd1\x06\xd2wQr\xa5\x14@^,\x14\x16\x8dv\xc2k\xfbr.\x17\xe5JÅ\x00\xd0W3U\xa3\xdc\xd8\xc9>\b\x81P\x9c<\x89c\xde\x11ED\x05\x01\x93V\xad\xa6sq\xd9\xc1\xf1\xbb\xf3\xb4\x8a\x12ա\xa7=\xa9,\x93\v\x05\xe7>\n\xa6\xfd\xd2rg\x95\xd6\x12ْ\xf7(\xb7\xbe\xea\x13ur\xaa\x11\xedz\x8fӖ\xefR\xe07ܗIÉI\xda\x05\xe5\x1c!):\x1a/RB\x12\xf16\xa2\x1d\x8al\xc6h#Prw)\x96\xab\xf3\x916\x1c\xacl\x84sD\x99\x8e\xc7P^B\xd7\x11\x14P`\x89G\\h\xe7h2\x830\xa6`\tw\xcdD\xa3p\xf0\xe1\x03h\v\x1fb\xcb\xff\xe1cL\xd7^H_\b5\xb1\x91\xd1\xf8,\xa4Lvߕ\xc5\x14\xbd\xb1\xcdн\xdfp\xc0?\x16\xe2\v?x\xea\x7f\xc2\u07bd\x86g&\xfcX\xee2\x98G\xd3\xee#T\xd8\x10\xc5Z\xf4\xbdUt\x12\xd0Z\xa2\x1c\x17T\xea\u07bfkSN1\xe3\x8e\xda\xdf\xddnl\xe7a\x14L\xecrw\x9b\xb8\xe51Da\xa4\x98A\x12\xbc\xce\x05\x94\xa0G\x0e\t\xc5\xe8}hC\x05\x1c/X[\x90\xe7\xd2\t\xb7\xb6\xa2\x15\xd4V\xa8q\xe6Ly'\xba\x90\xe5\x12Q\xb8\xb0?\xe4Л\b\x9c(\x86\xaak\x85\xc0EӠE\xe5c}\x8d\x86\x0f\x8f\xfbkw6\x92\xd3\xd9L0\x84\x0e\xabc\xc6 \xa7\xbe\x98\";8\xea].\xf2̶\xe8\x1f\xc366\xfc\xf3m\"\x9a\x9cC\x95\x9b.1T\b\x86\xe8F\x8dpx\xdcS\a\x96\xd9\xc6\xe1q\x8d\xf0r\x95\x83\xa1\xe1\xbd\x10\xc1\x15\xcaU\xfc\x06<\xaf9v\x83N\x01\xcc\xe9\r\x96\x0f\x8f\xb9B:\xba\x03\xfc\x91y\x92\x18.(P\xbdduB:\x1fC8\xbf\x0fo\xfd&\xc0\xfbW\x11\uf5d0/\xe0\xad^~5d*\xde\xc2\"_\xa3.^\x89\\\x01\xe6\x94\x1d\xac\xdf^\xa2\xf2\x96\x8b|w\xb5\x90YR\xfcb\xfaL\x96ˉ9\xd3,f\xa7G\xf2Mmh\xb8B\xbe\xb5\x11\x8d\x0fCC\xd8\xeb\xde\x06\x1a\x1a\x9e\x8b\xe8V\xf6]\xadh\x1d\x1fZ\xa6w\xea\xad\xf6m\xbd\"\xdc\xf7,\x9f\xd4;\x96\x12*^\xec\xd3kN\xae\x7f;\xeb\x8bK\x03=\x92:\xe4\x80'T@\xad6\x13\x12y\xd2\xe9J\xf8F\xddx\xb8\xf8\\/\xafH\xc1߃\xa2Pv\xa9gʀ^\xafK\xaf:t\xdd)H\xc5JB\xf5R\xb2J\xe2\x0e\xbc\xed/\xf5\x8fكҡs\xac\xdd\"\xea\xafQ*^\x15\x87%\xc0*j*\x96=\xed\xb5\x1bb\xff\xae\xa2\xa14\xdf\xc2p\xafy\x00\xa0\xbe\xe3\x95\xe4]XB\x0f\xbe\x01\xe6@2\xb9\x9c\x1f\xa1\xbd~=@\xd5w9f\xba\xc7\xe7\xcc\xe8\xe7\xbaF\x93c\xcb\x02\x0e\x16\r\xb3٩\xd5\xf3\xect2^vrę\xe6\xb2:\xc7\xf7\xcf\xcc\xdc\xdf\xc2ax\x97\xa7\a|[\xceN\xb7\xa3\xa3\x96\xe90\x87'J\xd5w\x15Z\xf2xx\x04M\xaeO,\x999\x80L\xf1Y\xc8\xce\x1aƞ0\xa8\xa2\x93LU*^\xe0R\x97̅3\x92\xe5\x8al\xdaɬ}9\x1f\x90D:#\xbd\xbf\xb7_\x19\x9f\x8c\xf3E8\xf7\ue6cb\xc2\xf4\x05w1?>\x05\xff6\x16^\xb9\xd0͟\xe6\xb7Z\xea\x99\xf0\x16\xc1\x0f\xff\n\xe4\xe8}\xca\xd4k^\x9e\x9b\xf9=)9\xeb\xa8\xd5`@\xce'\xba\x87g\x95\xe9H_\x8d\x8f\x85;\xf8\xcf\x7f\xaf\xfe\x17\x00\x00\xff\xff\x9f\xc23\x7f`\x1b\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcYK\x93۸\x11\xbeϯ\xe8r\x0e\xbe\x8c8\xebM*\x95\xd2\xcd\xd6d\xabTYOT\x96w\xee\x10\xd9$\xb1\x03\x02\f\x1e\x9aLR\xf9\xef\xa9\x06\b\x8a\x0fH\x94\x1c;<L\x8d\x00t\xa3\x1f读\x81\xd5ju\xc7Z\xfe\x8c\xdap%\xd7\xc0Z\x8e\xff\xb4(\xe9\x97\xc9^\xfeb2\xae\x1e\x8e\x1f\xee^\xb8,ְqƪ\xe6\v\x1a\xe5t\x8e\x8fXr\xc9-W\xf2\xaeA\xcb\nf\xd9\xfa\x0e\x80I\xa9,\xa3aC?\x01r%\xadVB\xa0^U(\xb3\x17w\xc0\x83\xe3\xa2@\xed\x99ǭ\x8f?e\x1f~\xce~\xba\x03\x90\xac\xc15\x10?\xd7\n\xc5\n\x93\x1dQ\xa0V\x19Ww\xa6Ŝ\xd8VZ\xb9v\r\xa7\x89@\xd6m\x19\xc4}d\x96\xfd\xe69\xf8A\xc1\x8d\xfd\xdbd\xe2Wn\xac\x9fl\x85\xd3L\x8cv\xf5\xe3\x86\xcb\xca\t\xa6\x873w\x00&W-\xaeቶlY\x8e4\xd6i\xe2EX\x01+\no\x1b&v\x9aK\x8bz\xa3\x84k\xa2MVP\xa0\xc95o\xad\xd7\xfd$\x10\x18ˬ3`\\^\x033\xf0\x84\xaf\x0f[\xb9Ӫ\xd2h\x82H\x00\xbf\x1b%w\xcc\xd6k\xc8\xc2\U000acb59\xc1n6\x98o\xef'\xba!\xfbF\xd2\x1a\xab\xb9\xacR\xfb\x7f\xe5\rB\xe1\xb4w\x1b\xe9\x9c#ؚ\x9b\xa1`\xaf̐p\xdabqV\f?ǑeM;\x95g@\x1a\x04*\x98Ŕ8\x1bմ\x02-\x16px\xb3\x18\x95(\x95n\x98]\x03\x97\xf6\xcf\x7f:o\x89\xceT\x99'}Trl\x96O4\n\x83\xe1 \ty\xa8B\x9d\xb4\x8d\xb2L\xfc/\x82Xb\xf0i@\x1f$\t|\x87㋢\xd0q\x03U\x82\xad\x11>\xb1\xfcŵ\xb0\xb7J\xb3\n\xe1W\x95\a\xe7\xbd֨;\xe7\x1d\xc2\x12S+'\n8D\x8d\x01\x8cU:\xe9\xc5\x16\xf3,Pu|#ۉ+\xc7{~\xe7C\x96kd\xc9C\x16Q&\xf3+\xb8\x92\xe9\x93\xf6\xb1«N\xd9КR\x15؛\x0e\x87\x12q\x03\xadV9\x1as\xe1\xdc\x13\xf9H\x86\xa7\xd3\xc0\xcc,a\xc5\xf1g&ښ}\b(\x93\xd7ذuG\xa1Z\x94\x1fw\xdb\xe7?\xeeG\xc3@\x82\xb4\xa8-\x8f0\x17\xbe\x01~\x0fFa\xac\xec{b\x18VAA\xc0\x8d\xc6kځ\x16\x16\x9d\f\xc1 ܀\xc6V\xa3Ai\x87.\x8e\x9f*\x81IP\x87\xdf1\xb7\x19\xecQ\x13\x9bx\xd0r%\x8f\xa8-h\xccU%\xf9\xbfz\xde\x06\xac\xf2\x9b\nf\xb1\xc3\xdd\xd3\xe7AR2\x01G&\x1c\xde\x03\x93\x054\xec\r4\xd2.\xe0䀟_b2\xf8\xac4\x02\x97\xa5ZCmmk\xd6\x0f\x0f\x15\xb71o\xe5\xaai\x9c\xe4\xf6\xed\xc1\xa7 ~pVi\xf3P\xe0\x11Ń\xe1Պ\xe9\xbc\xe6\x16s\xeb4>\xb0\x96\xaf\xbc\xe8\xd2箬)\xfe\xa0\xbbLgޏd\x9dy4|>\xe9\\\xf0\x00\xe5\x1e:N\xac#\rZ\x9c\fMCd\x9d/\x7f\xdd\x7f\x85\xb8\xb5w\xc6\xd4\xfa\xde\xee'Bsr\x01\x19\x8c\xcb\x12upb\xa9U\xe3y\xa2,Zť\xf5?r\xc1QN\xcdoܡ\xe1\x96\xfc\xfe\x0f\x87ƒ\xaf2\xd8\xf8d\x0e\a\x04\xd7R\x18\x15\x19l%lX\x83b\xc3\f\xfep\a\x90\xa5͊\f{\x9d\v\x86u\xc8tq\xb0\xda`\"\x96\x12g\xfcu\u0080}\x8b99\x8elGD\xbc\xe4\x1dؖJ\x03\x1b\xac\xccF\xec\xd2\xe1J_\x12c\xa7\x8b&\xf2|J\xd1D\xb1\xe4\x00\xcb\"쇕3\xa6\x00b\x9a+z\x1a\x8d\xad2\xdc*\xfdF\x8cC\x9a\xc8f\x1c\xce\x18\x9f\xbe\x9c\xc9\x1cł&\x1b\xbf\b\xb8,Ȏ؟9\x82\x87\xc0\xc0ˤd\xa5(&Ι7|[K4tD\rZ\xd2H&\x10\x9cK8\x15Q0,\x96\xa6Z\x1d\x94\x12Ȧx\x97\x1b\xbe\x97\xac5\xb5\xb2\v\xbamK\x88+\xbf\xbe\xb5H\x9bo\xf6\xdb{\xfa\x13\xc7\xe9\\\x1cy\xd1\x010\x05\x0f\x95\x13s\x90\x85\x00\xb4\xb4h\xb3߂\xe9\xc8\xe7F\x90N\bv\x10\xb8\x06\xab\xdd\\\xb1\xf3ǐ\xbe\xc8v#\x98I.\x98(\xb8\x1f\xaeO\x1d\xbf\xc8\x10r\xbf\xc2\xd6l\n5\xbd\xc5)\xffPU< \xe2}\xfe\x87Wn\xeb$\xe5\x85\xf3\a]u\xc3*\xbcZ\xa1\xc1\xf2\xa4>]\xb5\x13\xd4Q\xe5\x05ev\xcf\x1b\xaf\xef\x92f\x04\xcbߢY`y\xfe$\xcet{\x1e\x11\xa4\xb4\x9bHyN9E\x01F \x81\x05\xb8\xf6v\xd9)¹\xc6b.\xf3j\xe4\xaf\xc4\xf4X\xe93a;\x03wo\nf\xd9gu\xa4\xd6K\x96\xbc\x9a\xef=\xec\xd1.\xc5\xc8E\xd5fIc\xb0%Y\x9cr\x04I\xb2jh|\x15\x13\b\x95I%\xaf\xbar8\xb1i\xc9Q\x14\xe6\xe6h_\xb0\x87\x17b\x01\xc3z%b\xb6렊\xe8\xc13\xe8\x0e\x843\xbeE\xa3Ʉ\x02!\xa7d\x04\x89'\x8e\xdc\xc0\xbbw\xa04\xbc\v\xad\xfb\xbb\xfb\x90\x7f\x1c\x17v\xc5\xe5`\x8f\x04\xc7W.D\xdc\xf7\xa6\xb4D\xce\xed;\x06\xe5\x96@\xfc\xef\x93\xe5\x13;Xjd\xbc\xeeV\xc1+㶯]S\b\x1ey\x99{8`I\xf5\x92F봤ԆZS\x05a<K\xe5\x12\xd8~A)3\xc83\v\nMS\x92ׂ\xfe\x9fb\xf60\xd0\x13ʸ\xf66\t}\t\xdb_\x92,\t9^\x1d\xe5T\x9aW\x9c\xfa\x02\xd9Ϝ\xea\x96\x00\x0e\tI\xbb\xf6\xd7Õ\xc7ی*\x85X\xc4\x11\x00\x9e\xd8Q\x84\x86\xcd\t\xc0\xa9\xef\xd8\xec\xb7\t\x9e=E\xd1\xc5W\":\x17\xad\xb1{\xde\\e\a\x12%\x81\xd74\xfcZ\xf3\xbc\x1e\xfbm\xd6#xY\xd8\v\xfa\x1a\xf5&1\xc3\x1d\xd7y\xd0\x1c\xc9\xfa\xdbh\xf1$NF\b\xe7\xadL\xa3\x91\xff\\\xa8˅J\xcb4\x13\x02\xc5/\\\xa0\t\xfb^\x91\x01ws\xaaެ\xae9\xa0&Ö4\xd9op&\av\xb7\x82\x14\"-j*\xdaB\xb5\xe9Ll\xdc\xcek\x06\xc9\u06dd\xf9l\x02\xb6әs\x95n!&k\xa6\xb07\x99\x1e\x02\xc8tj\x1c\x8b\xc9\xd9\xdd\xf3\xe6\xaa6\xcb_\x91\\\xd7h\x85\xbb\xcf\xce?\xb9\xd3\x1a\xa5\x8d7\xa2\xaa\xfc\xa6V+\x0fw\x89\xc3ۢ\xa5\xf6dN\xe1\xef2t1\x80\x7f\x16[&\x7fc\x15\xef+S\x9e?\xb1\v\x94\xfen\x85\xb8a\x01xD\t\x14\x1b\x8c\vJ\xa5\x9e\xa5ɦ4i|\xeb\xb9ti%\x1c\xbfx\x18;\xf1\xe2\x1d\xcdWB\v\x7fI\xf0\xde\\\xe0\xe9\xb3\x1aEj\xc2\bs\x88\x89\x17\xa1\x05\xb3\xb8J2\xbd\xaaXI\xc2P_\xbc}A\xe3D\"c\xff\xc0\xe2-l\x19\xfa_\x93,\xde.wm\xcc\x00\x03\x1d\x98t\xb8}\xa9\x89\xfd\xf6\x8a\xaeAcX\xb5\x94X?\x87U\xe1\xee\xa9#\x01v\xa0\xc2f,\xda{\xd3\x05\xdbM\tC\xaabI\x82'U\xf8\xed\xe5\xcdW\xae7I\xd22[/H\xb2c\xb6\x8e\x00S:!<ͬ\f꺢\x03R4}\xafj\xc8_;,\x89GkR\x00\x88\xd7\x1c$\x94\xaeIuXO\xf8\x9a\x18\xfd\x98\xe7\xd8ڄf+\xd8il\x99NN\xcd\x1e\xa3\x86\x93\xe1f'\x15hq.ɳ\x7f\xefI\xcc\xfd\xe2\xa1\xf1&;w\xf2-\x99:^\a\xd5JDd\xf7O2\xa7\xba\xc0?\xfa\x8c;\xe5T\xb4\xcbb\xe4\xae\x01}_\x81zN\x84\xc2\xdc\xc4\xeb\xaa\xd8@\x14ܴ\x82\xbd\xa5*\x9e(\xe1\x00m\x06q\x1b\x11>&\xf9[\v\xaa\xfe\x81,]F\xa5^\xb9R>8W\xd1\x00\x9c\x1e\xbe~\xcc\x0e\x17\x801F\xf2\xf6\xf1\xca\xceh\xfb\x18\xa3\x8e\x17(-\x15\xb1\xba/YOe\xb6\xbc\xd8\xeb\x0e\xeeOo\xeb\fFϦK\x12\x8f\x16/T&݃m\xaa.\xd9S\x88\x13\xb0\xf8G\x82\xcd\xf4I\xed\xbe\x7f\xa1c\xb6{\xa9\xc8k&+\n\bI\xc9\xcd'\xc7\x14\xe3Y\xa91*,\xc6\xe2\xff?k\x8a\xe4q\x99\rzɋ\x01\xef\xeezj8\xe2\x0e\xfd\xcb\xd0\x1a\xfe\xfd\x9f\xbb\xff\x06\x00\x00\xff\xff\xf1\x97\xa7\xf7F!\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := io.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1.CustomResourceDefinition))
	}
	return objs
}
