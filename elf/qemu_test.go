package elf

var qemufiles = map[string]ef{
	"/usr/bin/qemu-system-x86_64": {[]string{
		"libvirglrenderer.so.0",
		"libepoxy.so.0",
		"libdrm.so.2",
		"libgbm.so.1",
		"libX11.so.6",
		"libz.so.1",
		"libaio.so.1",
		"libpixman-1.so.0",
		"libutil.so.1",
		"libnuma.so.1",
		"libbluetooth.so.3",
		"libncursesw.so.6",
		"libbrlapi.so.0.6",
		"libasound.so.2",
		"libpulse.so.0",
		"libvdeplug.so.3",
		"libpng16.so.16",
		"libjpeg.so.8",
		"libsasl2.so.3",
		"libSDL2-2.0.so.0",
		"libvte-2.91.so.0",
		"libgtk-3.so.0",
		"libgdk-3.so.0",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libgnutls.so.30",
		"libnettle.so.6",
		"liblzo2.so.2",
		"libsnappy.so.1",
		"libseccomp.so.2",
		"libspice-server.so.1",
		"libcacard.so.0",
		"libusb-1.0.so.0",
		"libusbredirparser.so.1",
		"libjemalloc.so.2",
		"libgmodule-2.0.so.0",
		"librt.so.1",
		"libm.so.6",
		"libgcc_s.so.1",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libvirglrenderer.so.0": {[]string{
		"libm.so.6",
		"libgbm.so.1",
		"libepoxy.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libm.so.6": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libc.so.6": {[]string{
		"ld-linux-x86-64.so.2"}},

	"/lib64/ld-linux-x86-64.so.2": {},

	"/lib64/libgbm.so.1": {[]string{
		"libexpat.so.1",
		"libm.so.6",
		"libdl.so.2",
		"libwayland-client.so.0",
		"libwayland-server.so.0",
		"libdrm.so.2",
		"libc.so.6"}},

	"/lib64/libexpat.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libdl.so.2": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libwayland-client.so.0": {[]string{
		"libffi.so.6",
		"librt.so.1",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libffi.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/librt.so.1": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libpthread.so.0": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libwayland-server.so.0": {[]string{
		"libffi.so.6",
		"librt.so.1",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libdrm.so.2": {[]string{
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libepoxy.so.0": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/lib64/libX11.so.6": {[]string{
		"libxcb.so.1",
		"libdl.so.2",
		"libc.so.6"}},

	"/lib64/libxcb.so.1": {[]string{
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/lib64/libXau.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/libXdmcp.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/libz.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libaio.so.1": {},

	"/lib64/libpixman-1.so.0": {[]string{
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libutil.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libnuma.so.1": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libbluetooth.so.3": {[]string{
		"libc.so.6"}},

	"/lib64/libncursesw.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/libbrlapi.so.0.6": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libasound.so.2": {[]string{
		"libm.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"librt.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libpulse.so.0": {[]string{
		"libpulsecommon-10.0.so",
		"libdbus-1.so.3",
		"libpthread.so.0",
		"libdl.so.2",
		"libm.so.6",
		"libc.so.6"},
		{},
		[]string{"/lib64/pulseaudio"}},

	"/lib64/pulseaudio/libpulsecommon-10.0.so": {[]string{
		"libxcb.so.1",
		"libsystemd.so.0",
		"libsndfile.so.1",
		"libasyncns.so.0",
		"libdbus-1.so.3",
		"libpthread.so.0",
		"librt.so.1",
		"libdl.so.2",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libsystemd.so.0": {[]string{
		"libresolv.so.2",
		"librt.so.1",
		"liblzma.so.5",
		"liblz4.so.1",
		"libgcrypt.so.20",
		"libgpg-error.so.0",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libresolv.so.2": {[]string{
		"libc.so.6"}},

	"/lib64/liblzma.so.5": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/liblz4.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libgcrypt.so.20": {[]string{
		"libgpg-error.so.0",
		"libc.so.6"}},

	"/lib64/libgpg-error.so.0": {[]string{
		"libc.so.6"}},

	"/lib64/libsndfile.so.1": {[]string{
		"libFLAC.so.8",
		"libogg.so.0",
		"libvorbis.so.0",
		"libvorbisenc.so.2",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libFLAC.so.8": {[]string{
		"libogg.so.0",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libogg.so.0": {[]string{
		"libc.so.6"}},

	"/lib64/libvorbis.so.0": {[]string{
		"libm.so.6",
		"libogg.so.0",
		"libc.so.6"}},

	"/lib64/libvorbisenc.so.2": {[]string{
		"libvorbis.so.0",
		"libm.so.6",
		"libogg.so.0",
		"libc.so.6"}},

	"/lib64/libasyncns.so.0": {[]string{
		"libresolv.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libdbus-1.so.3": {[]string{
		"libpthread.so.0",
		"libsystemd.so.0",
		"libc.so.6"}},

	"/lib64/libvdeplug.so.3": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/lib64/libpng16.so.16": {[]string{
		"libz.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libjpeg.so.8": {[]string{
		"libc.so.6"}},

	"/lib64/libsasl2.so.3": {[]string{
		"libdl.so.2",
		"libresolv.so.2",
		"libc.so.6"}},

	"/lib64/libSDL2-2.0.so.0": {[]string{
		"libm.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"librt.so.1",
		"libc.so.6"}},

	"/lib64/libvte-2.91.so.0": {[]string{
		"libgtk-3.so.0",
		"libgdk-3.so.0",
		"libpangocairo-1.0.so.0",
		"libpango-1.0.so.0",
		"libatk-1.0.so.0",
		"libcairo-gobject.so.2",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libz.so.1",
		"libpcre2-8.so.0",
		"libgnutls.so.30",
		"libstdc++.so.6",
		"libm.so.6",
		"libc.so.6",
		"libgcc_s.so.1"}},

	"/lib64/libgtk-3.so.0": {[]string{
		"libgdk-3.so.0",
		"libgmodule-2.0.so.0",
		"libpangocairo-1.0.so.0",
		"libX11.so.6",
		"libXi.so.6",
		"libXfixes.so.3",
		"libcairo-gobject.so.2",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libatk-1.0.so.0",
		"libatk-bridge-2.0.so.0",
		"libepoxy.so.0",
		"libpangoft2-1.0.so.0",
		"libpango-1.0.so.0",
		"libfontconfig.so.1",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libgdk-3.so.0": {[]string{
		"libpangocairo-1.0.so.0",
		"libpango-1.0.so.0",
		"libgdk_pixbuf-2.0.so.0",
		"libcairo-gobject.so.2",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libfontconfig.so.1",
		"libXinerama.so.1",
		"libXi.so.6",
		"libXrandr.so.2",
		"libXcursor.so.1",
		"libXcomposite.so.1",
		"libXdamage.so.1",
		"libXfixes.so.3",
		"libxkbcommon.so.0",
		"libwayland-cursor.so.0",
		"libwayland-egl.so.1",
		"libwayland-client.so.0",
		"libX11.so.6",
		"libXext.so.6",
		"libcairo.so.2",
		"libepoxy.so.0",
		"libm.so.6",
		"librt.so.1",
		"libc.so.6"}},

	"/lib64/libpangocairo-1.0.so.0": {[]string{
		"libpango-1.0.so.0",
		"libpangoft2-1.0.so.0",
		"libm.so.6",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libcairo.so.2",
		"libc.so.6"}},

	"/lib64/libpango-1.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libm.so.6",
		"libthai.so.0",
		"libc.so.6"}},

	"/lib64/libgobject-2.0.so.0": {[]string{
		"libglib-2.0.so.0",
		"libffi.so.6",
		"libc.so.6"}},

	"/lib64/libglib-2.0.so.0": {[]string{
		"libpcre.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libpcre.so.1": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libthai.so.0": {[]string{
		"libdatrie.so.1",
		"libc.so.6"}},

	"/lib64/libdatrie.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libpangoft2-1.0.so.0": {[]string{
		"libpango-1.0.so.0",
		"libm.so.6",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libharfbuzz.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libc.so.6"}},

	"/lib64/libharfbuzz.so.0": {[]string{
		"libglib-2.0.so.0",
		"libfreetype.so.6",
		"libgraphite2.so.3",
		"libc.so.6"}},

	"/lib64/libfreetype.so.6": {[]string{
		"libz.so.1",
		"libbz2.so.1.0",
		"libpng16.so.16",
		"libharfbuzz.so.0",
		"libc.so.6"}},

	"/lib64/libbz2.so.1.0": {[]string{
		"libc.so.6"}},

	"/lib64/libgraphite2.so.3": {[]string{
		"libc.so.6"}},

	"/lib64/libfontconfig.so.1": {[]string{
		"libfreetype.so.6",
		"libexpat.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libcairo.so.2": {[]string{
		"libpthread.so.0",
		"libpixman-1.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libEGL.so.1",
		"libdl.so.2",
		"libpng16.so.16",
		"libxcb-shm.so.0",
		"libxcb.so.1",
		"libxcb-render.so.0",
		"libXrender.so.1",
		"libX11.so.6",
		"libXext.so.6",
		"libz.so.1",
		"libGL.so.1",
		"librt.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libEGL.so.1": {[]string{
		"libm.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libGLdispatch.so.0": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/lib64/libxcb-shm.so.0": {[]string{
		"libxcb.so.1",
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/lib64/libxcb-render.so.0": {[]string{
		"libxcb.so.1",
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/lib64/libXrender.so.1": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libXext.so.6": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libGL.so.1": {[]string{
		"libGLX.so.0",
		"libX11.so.6",
		"libXext.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libGLX.so.0": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libgdk_pixbuf-2.0.so.0": {[]string{
		"libgmodule-2.0.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libpng16.so.16",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libgmodule-2.0.so.0": {[]string{
		"libdl.so.2",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libgio-2.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libgmodule-2.0.so.0",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libz.so.1",
		"libresolv.so.2",
		"libmount.so.1",
		"libc.so.6"}},

	"/lib64/libmount.so.1": {[]string{
		"libblkid.so.1",
		"libuuid.so.1",
		"librt.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libblkid.so.1": {[]string{
		"libuuid.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libuuid.so.1": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libcairo-gobject.so.2": {[]string{
		"libcairo.so.2",
		"libpthread.so.0",
		"libpixman-1.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libEGL.so.1",
		"libdl.so.2",
		"libpng16.so.16",
		"libxcb-shm.so.0",
		"libxcb.so.1",
		"libxcb-render.so.0",
		"libXrender.so.1",
		"libX11.so.6",
		"libXext.so.6",
		"libz.so.1",
		"libGL.so.1",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"librt.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libXinerama.so.1": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libc.so.6"}},

	"/lib64/libXi.so.6": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libc.so.6"}},

	"/lib64/libXrandr.so.2": {[]string{
		"libXext.so.6",
		"libXrender.so.1",
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libXcursor.so.1": {[]string{
		"libXrender.so.1",
		"libXfixes.so.3",
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libXfixes.so.3": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libXcomposite.so.1": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libXdamage.so.1": {[]string{
		"libXfixes.so.3",
		"libX11.so.6",
		"libc.so.6"}},

	"/lib64/libxkbcommon.so.0": {[]string{
		"libc.so.6"}},

	"/lib64/libwayland-cursor.so.0": {[]string{
		"libwayland-client.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libwayland-egl.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libatk-1.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libc.so.6"}},

	"/lib64/libatk-bridge-2.0.so.0": {[]string{
		"libatk-1.0.so.0",
		"libgobject-2.0.so.0",
		"libatspi.so.0",
		"libdbus-1.so.3",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libatspi.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libX11.so.6",
		"libdbus-1.so.3",
		"libglib-2.0.so.0",
		"libc.so.6"}},

	"/lib64/libpcre2-8.so.0": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libgnutls.so.30": {[]string{
		"libz.so.1",
		"libp11-kit.so.0",
		"libunistring.so.2",
		"libtasn1.so.6",
		"libnettle.so.6",
		"libhogweed.so.4",
		"libgmp.so.10",
		"libc.so.6"}},

	"/lib64/libp11-kit.so.0": {[]string{
		"libffi.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libunistring.so.2": {[]string{
		"libc.so.6"}},

	"/lib64/libtasn1.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/libnettle.so.6": {[]string{
		"libc.so.6"}},

	"/lib64/libhogweed.so.4": {[]string{
		"libnettle.so.6",
		"libgmp.so.10",
		"libc.so.6"}},

	"/lib64/libgmp.so.10": {[]string{
		"libc.so.6"}},

	"/lib64/libstdc++.so.6": {[]string{
		"libm.so.6",
		"libc.so.6",
		"ld-linux-x86-64.so.2",
		"libgcc_s.so.1"}},

	"/lib64/libgcc_s.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/liblzo2.so.2": {[]string{
		"libc.so.6"}},

	"/lib64/libsnappy.so.1": {[]string{
		"libstdc++.so.6",
		"libm.so.6",
		"libc.so.6",
		"libgcc_s.so.1"}},

	"/lib64/libseccomp.so.2": {[]string{
		"libc.so.6"}},

	"/lib64/libspice-server.so.1": {[]string{
		"libcelt051.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libjpeg.so.8",
		"libpixman-1.so.0",
		"libsasl2.so.3",
		"libssl.so.1.1",
		"libcrypto.so.1.1",
		"libz.so.1",
		"libm.so.6",
		"librt.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libcelt051.so.0": {[]string{
		"libm.so.6",
		"libc.so.6"}},

	"/lib64/libssl.so.1.1": {[]string{
		"libcrypto.so.1.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libcrypto.so.1.1": {[]string{
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libcacard.so.0": {[]string{
		"libglib-2.0.so.0",
		"libssl3.so",
		"libsmime3.so",
		"libnss3.so",
		"libnssutil3.so",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so",
		"libc.so.6"}},

	"/lib64/libssl3.so": {[]string{
		"libnss3.so",
		"libnssutil3.so",
		"libpthread.so.0",
		"libc.so.6",
		"libz.so.1",
		"libplc4.so",
		"libnspr4.so"}},

	"/lib64/libnss3.so": {[]string{
		"libnssutil3.so",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/lib64/libnssutil3.so": {[]string{
		"libpthread.so.0",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/lib64/libplds4.so": {[]string{
		"libnspr4.so",
		"libc.so.6"}},

	"/lib64/libnspr4.so": {[]string{
		"libpthread.so.0",
		"libdl.so.2",
		"librt.so.1",
		"libc.so.6"}},

	"/lib64/libplc4.so": {[]string{
		"libnspr4.so",
		"libc.so.6"}},

	"/lib64/libsmime3.so": {[]string{
		"libnss3.so",
		"libnssutil3.so",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/lib64/libusb-1.0.so.0": {[]string{
		"libudev.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/lib64/libudev.so.1": {[]string{
		"libresolv.so.2",
		"librt.so.1",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/lib64/libusbredirparser.so.1": {[]string{
		"libc.so.6"}},

	"/lib64/libjemalloc.so.2": {[]string{
		"libstdc++.so.6",
		"libpthread.so.0",
		"libdl.so.2",
		"libgcc_s.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},
}

var qemuresolved = []string{
	"/lib64/libgpg-error.so.0",
	"/lib64/libxcb-shm.so.0",
	"/lib64/libcrypto.so.1.1",
	"/lib64/libvdeplug.so.3",
	"/lib64/libEGL.so.1",
	"/lib64/libunistring.so.2",
	"/lib64/libnettle.so.6",
	"/lib64/ld-linux-x86-64.so.2",
	"/lib64/libpng16.so.16",
	"/lib64/libglib-2.0.so.0",
	"/lib64/libgmodule-2.0.so.0",
	"/lib64/libXi.so.6",
	"/lib64/libmount.so.1",
	"/lib64/libnss3.so",
	"/lib64/libdl.so.2",
	"/lib64/libdrm.so.2",
	"/lib64/libFLAC.so.8",
	"/lib64/libthai.so.0",
	"/lib64/libbz2.so.1.0",
	"/lib64/libcairo-gobject.so.2",
	"/lib64/libwayland-server.so.0",
	"/lib64/libXext.so.6",
	"/lib64/libblkid.so.1",
	"/lib64/libatspi.so.0",
	"/lib64/libasound.so.2",
	"/lib64/libpango-1.0.so.0",
	"/lib64/libXrender.so.1",
	"/lib64/libgio-2.0.so.0",
	"/lib64/libatk-1.0.so.0",
	"/lib64/libsnappy.so.1",
	"/lib64/libssl.so.1.1",
	"/lib64/libXau.so.6",
	"/lib64/libfontconfig.so.1",
	"/lib64/libXcomposite.so.1",
	"/lib64/libXdamage.so.1",
	"/lib64/libwayland-cursor.so.0",
	"/lib64/libGL.so.1",
	"/lib64/libuuid.so.1",
	"/lib64/libwayland-egl.so.1",
	"/lib64/libsndfile.so.1",
	"/lib64/libvorbis.so.0",
	"/lib64/libdatrie.so.1",
	"/lib64/libGLdispatch.so.0",
	"/lib64/libxcb-render.so.0",
	"/lib64/libplc4.so",
	"/lib64/libharfbuzz.so.0",
	"/lib64/libplds4.so",
	"/lib64/libpixman-1.so.0",
	"/lib64/libjpeg.so.8",
	"/lib64/libgcc_s.so.1",
	"/lib64/liblz4.so.1",
	"/lib64/libogg.so.0",
	"/lib64/libGLX.so.0",
	"/lib64/libXrandr.so.2",
	"/lib64/libXcursor.so.1",
	"/lib64/libz.so.1",
	"/lib64/libbluetooth.so.3",
	"/lib64/liblzma.so.5",
	"/lib64/libsasl2.so.3",
	"/lib64/libpangocairo-1.0.so.0",
	"/lib64/libudev.so.1",
	"/lib64/libm.so.6",
	"/lib64/libaio.so.1",
	"/lib64/libfreetype.so.6",
	"/lib64/libatk-bridge-2.0.so.0",
	"/lib64/libtasn1.so.6",
	"/lib64/liblzo2.so.2",
	"/lib64/libXdmcp.so.6",
	"/lib64/libgdk_pixbuf-2.0.so.0",
	"/lib64/pulseaudio/libpulsecommon-10.0.so",
	"/lib64/libsystemd.so.0",
	"/lib64/libhogweed.so.4",
	"/lib64/libssl3.so",
	"/lib64/libgraphite2.so.3",
	"/lib64/libpcre2-8.so.0",
	"/lib64/libgmp.so.10",
	"/lib64/libwayland-client.so.0",
	"/lib64/libffi.so.6",
	"/lib64/librt.so.1",
	"/lib64/libutil.so.1",
	"/lib64/libnuma.so.1",
	"/lib64/libjemalloc.so.2",
	"/lib64/libcairo.so.2",
	"/lib64/libstdc++.so.6",
	"/lib64/libusbredirparser.so.1",
	"/lib64/libc.so.6",
	"/lib64/libncursesw.so.6",
	"/lib64/libbrlapi.so.0.6",
	"/lib64/libasyncns.so.0",
	"/lib64/libgobject-2.0.so.0",
	"/lib64/libgcrypt.so.20",
	"/lib64/libp11-kit.so.0",
	"/lib64/libspice-server.so.1",
	"/lib64/libpulse.so.0",
	"/lib64/libgnutls.so.30",
	"/lib64/libseccomp.so.2",
	"/lib64/libXfixes.so.3",
	"/usr/bin/qemu-system-x86_64",
	"/lib64/libvirglrenderer.so.0",
	"/lib64/libexpat.so.1",
	"/lib64/libpthread.so.0",
	"/lib64/libresolv.so.2",
	"/lib64/libcelt051.so.0",
	"/lib64/libxcb.so.1",
	"/lib64/libvorbisenc.so.2",
	"/lib64/libgdk-3.so.0",
	"/lib64/libpcre.so.1",
	"/lib64/libXinerama.so.1",
	"/lib64/libdbus-1.so.3",
	"/lib64/libSDL2-2.0.so.0",
	"/lib64/libcacard.so.0",
	"/lib64/libnssutil3.so",
	"/lib64/libgbm.so.1",
	"/lib64/libsmime3.so",
	"/lib64/libxkbcommon.so.0",
	"/lib64/libusb-1.0.so.0",
	"/lib64/libgtk-3.so.0",
	"/lib64/libepoxy.so.0",
	"/lib64/libX11.so.6",
	"/lib64/libvte-2.91.so.0",
	"/lib64/libpangoft2-1.0.so.0",
	"/lib64/libnspr4.so",
}
