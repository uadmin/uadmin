//
//jRoll - https://fitsbach.github.io/jRoll/
//Version - 0.1.6
//Licensed unter the GNU General Public License - gnu.org/licenses/gpl.html
//
//Copyright (c) 2016 Jimmy Fitzback
//

(function($) {

    $.fn.jRoll = function(options) {

        // Default options
        var settings = $.extend({
            radius: 100,
            animation: "heartbeat",
			colors: ['#003056','#04518C','#00A1D9','#47D9BF','#F2D03B'],
			monocolor: false
        }, options );

		//Fill the colors array if it's not full(3 colors)
		switch(settings.colors.length){
			case 0:
				settings.colors = ['#003056','#04518C','#00A1D9','#47D9BF','#F2D03B'];
				break;
			case 1:
				settings.colors[1]='#04518C';
				settings.colors[2]='#00A1D9';
				break;
			case 2:
				settings.colors[2]='#00A1D9';
				settings.colors[3]='#47D9BF';
				settings.colors[4]='#F2D03B';
				break;
			case 3:
				settings.colors[3]='#47D9BF';
				settings.colors[4]='#F2D03B';
				break;
			case 4:
				settings.colors[4]='#F2D03B';
				break;
		}
		if(settings.monocolor==true){
			settings.colors[1]=settings.colors[0];
			settings.colors[2]=settings.colors[0];
			settings.colors[3]=settings.colors[0];
			settings.colors[4]=settings.colors[0];
		}
        switch(settings.animation){
			case 'heartbeat':
			//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");;
				//Rings CSS
				var inRingStyle = "animation: heartbeatIn 1s linear 0s infinite;";
				var midRingStyle = "animation: heartbeatMid 1s linear 0.3s infinite;";
				var outRingStyle = "animation: heartbeatOut 1s linear 0.315s infinite;";

				//Rings Sizes
				var inRingSize = settings.radius/4;
				var midRingSize = settings.radius/3 ;
				var halfradius = settings.radius/2;
				var inRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollInRing" style="'+inRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+inRingSize+'" stroke="'+settings.colors[0]+'" stroke-width="6" fill="'+settings.colors[4]+'"></circle></svg>');
				var midRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollMidRing" style="'+midRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+midRingSize+'" stroke="'+settings.colors[1]+'" stroke-width="3" fill="transparent"></circle></svg>');
				var outRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollOutRing" style="'+outRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+((settings.radius/2)-0)+'" stroke="'+settings.colors[2]+'" stroke-width="4" fill="transparent"></circle></svg>');
				$(this).append(inRing).append(midRing).append(outRing);
				break;

			case 'pulse':
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");;
				//Rings CSS
				var inRingStyle = "animation: pulseIn 1s linear 0s infinite;";
				var midRingStyle = "animation: pulseMid 1s linear 0s infinite;";
				var outRingStyle = "animation: pulseOut 1s linear 0s infinite;";

				//Rings Sizes
				var inRingSize = settings.radius/4;
				var midRingSize = settings.radius/3 ;
				var halfradius = settings.radius/2;
				var strokew = settings.radius/12;
				var inRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+inRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+inRingSize+'" stroke="'+settings.colors[0]+'" stroke-width="'+strokew+'" fill="#64d4ce"></circle></svg>');
				var midRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+midRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+midRingSize+'" stroke="'+settings.colors[1]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				var outRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+outRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+halfradius+'" stroke="'+settings.colors[2]+'" stroke-width="2" fill="transparent"></circle></svg>');
				$(this).append(inRing).append(midRing).append(outRing);
				break;

			case 'slicedspinner':
				//Parent CSS
				$(this).css("width", (settings.radius*2)+'px').css("height",(settings.radius*2)+'px').css("overflow","hidden");
				//Animation has to be runned on parent container because of the Circle hack.
				$(this).css("animation", "slicedspinner 1s linear 0s infinite" );

				//SVG
				var Ring1= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[0]+'" stroke-width="2" fill="transparent"></circle></svg>');
				var Ring2= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" style="left:'+settings.radius+'px; transform: rotate(90deg);" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[1]+'" stroke-width="2" fill="transparent"></circle></svg>');
				var Ring3= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" style="left:'+settings.radius+'px; top:'+settings.radius+'px; transform: rotate(180deg);" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[0]+'" stroke-width="2" fill="transparent"></svg>');
				var Ring4= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" style="top:'+settings.radius+'px; transform: rotate(270deg);" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[1]+'" stroke-width="2" fill="transparent"></svg>');
				$(this).append(Ring1).append(Ring2).append(Ring3).append(Ring4);
				break;

			case 'halfslicedspinner':
				//Parent CSS
				$(this).css("width", (settings.radius*2)+'px').css("height",(settings.radius*2)+'px').css("overflow","hidden");
				//Animation has to be runned on parent container because of the Circle hack.
				$(this).css("animation", "slicedspinner 1s linear 0s infinite" );

				//SVG
				var Ring1= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[0]+'" stroke-width="2" fill="transparent"></circle></svg>');
				var Ring2= $('<svg height="'+settings.radius+'" width="'+settings.radius+'" style="left:'+settings.radius+'px; top:'+settings.radius+'px; transform: rotate(180deg);" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-2)+'" stroke="'+settings.colors[0]+'" stroke-width="2" fill="transparent"></svg>');
				$(this).append(Ring1).append(Ring2);
				break;

			case 'gyroscope':
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden").css("animation","gyroscope3D 2s linear 0s infinite");
				//Rings CSS
				var inRingStyle = "animation: gyroscopeIn 2s linear 0s infinite;z-index:1;";
				var outRingStyle = "animation: gyroscopeOut 2s linear 0s infinite;z-index:2;";

				//Rings Sizes
				var inRingSize = settings.radius/4;
				var halfradius = settings.radius/2;
				var strokew = settings.radius/12;
				var inRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+inRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-(strokew*4)-strokew)+'" stroke="'+settings.colors[0]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				var outRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+outRingStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius-(strokew*4))+'" stroke="'+settings.colors[1]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				//var midBall= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="z-index:-2;" ><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+(settings.radius/2)+'" fill="'+settings.colors[1]+'"></circle></svg>');

				$(this).append(inRing).append(outRing);
				break;

			case 'wave':
				if(settings.colors.length<=3){
					settings.colors[3]= '#DB9E36';
				}
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Wave1Style = "animation: waveOut 1.5s linear 0s infinite;";
				var Wave2Style = "animation: waveMid 1.5s linear 0s infinite;";
				var Wave3Style = "animation: waveIn 1.5s linear 0s infinite;";
				var WaveCenterStyle = "animation: waveCenter 1.5s linear 0s infinite;";

				//Rings Sizes
				var strokew = settings.radius/12;

				var Wave1= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+Wave1Style+'"><circle cx="'+settings.radius+'" cy="'+(settings.radius*2)+'" r="'+(settings.radius-(strokew*2))+'" stroke="'+settings.colors[0]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				var Wave2= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+Wave2Style+'"><circle cx="'+settings.radius+'" cy="'+(settings.radius*2)+'" r="'+(settings.radius-(strokew*4))+'" stroke="'+settings.colors[1]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				var Wave3= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+Wave3Style+'"><circle cx="'+settings.radius+'" cy="'+(settings.radius*2)+'" r="'+(settings.radius-strokew*6)+'" stroke="'+settings.colors[2]+'" stroke-width="'+strokew+'" fill="transparent"></circle></svg>');
				var WaveCenter= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" style="'+WaveCenterStyle+'"><circle cx="'+settings.radius+'" cy="'+(settings.radius*2)+'" r="'+(settings.radius-(strokew*8))+'" stroke="'+settings.colors[3]+'" stroke-width="'+strokew+'" fill="'+settings.colors[0]+'"></circle></svg>');

				$(this).append(Wave1).append(Wave2).append(Wave3).append(WaveCenter);
				break;

			case 'jumpdots':
				if(settings.colors.length<=3){
					settings.colors[3]= '#DB9E36';
					settings.colors[4] = '#BD4932';
				}
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden").css("animation","jumpdotdiv 2s linear 1s infinite");

				//Waves CSS
				var Dot1Style = "animation: jumpdots 2s linear 0s infinite;margin-left:calc(50% - "+((settings.radius/12)*8)+"px);margin-top:"+settings.radius+"px;";
				var Dot2Style = "animation: jumpdots 2s linear 0.2s infinite;margin-left:calc(50% - "+((settings.radius/12)*4)+"px);margin-top:"+settings.radius+"px;";
				var Dot3Style = "animation: jumpdots 2s linear 0.4s infinite;margin-left:calc(50%);margin-top:"+settings.radius+"px;";
				var Dot4Style = "animation: jumpdots 2s linear 0.6s infinite;margin-left:calc(50% - "+((settings.radius/12)*-4)+"px);margin-top:"+settings.radius+"px;";
				var Dot5Style = "animation: jumpdots 2s linear 0.8s infinite;margin-left:calc(50% - "+((settings.radius/12)*-8)+"px);margin-top:"+settings.radius+"px;";


				var Dot1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[3]+'"></circle></svg>');
				var Dot5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[4]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(Dot5);
				break;

			case '3dots':
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: suspensionpoint 3s linear -2s infinite;margin-left:calc(50% - "+(((settings.radius/4)/2)*4)+"px);margin-top: calc(50% - "+(((settings.radius/4)/2))+"px);";
				var Dot2Style = "animation: suspensionpoint 3s linear -1s infinite;margin-left:calc(50% - "+(((settings.radius/4)/2)*4)+"px);margin-top: calc(50% - "+(((settings.radius/4)/2))+"px);";
				var Dot3Style = "animation: suspensionpoint 3s linear 0s infinite;margin-left:calc(50% - "+(((settings.radius/4)/2)*4)+"px);margin-top: calc(50% - "+(((settings.radius/4)/2))+"px);";


				var Dot1= $('<svg height="'+((settings.radius/4))+'" width="'+((settings.radius/4))+'" style="'+Dot1Style+'"><circle style="transform-origin: '+(settings.radius/8)+'px '+(settings.radius/8)+'px; animation: suspensionpointcircle 3s linear -2s infinite" cx="50%" cy="50%" r="'+(settings.radius/8)+'" fill="rgba(0,0,0,0)" stroke="'+settings.colors[3]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/4))+'" width="'+((settings.radius/4))+'" style="'+Dot2Style+'"><circle style="transform-origin: '+(settings.radius/8)+'px '+(settings.radius/8)+'px; animation: suspensionpointcircle 3s linear -1s infinite" cx="50%" cy="50%" r="'+(settings.radius/8)+'" fill="rgba(0,0,0,0)" stroke="'+settings.colors[2]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/4))+'" width="'+((settings.radius/4))+'" style="'+Dot3Style+'"><circle style="transform-origin: '+(settings.radius/8)+'px '+(settings.radius/8)+'px; animation: suspensionpointcircle 3s linear 0s infinite" cx="50%" cy="50%" r="'+(settings.radius/8)+'" fill="rgba(0,0,0,0)" stroke="'+settings.colors[1]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3);
				break;

			case 'popdot':
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				//var Dot1Style = "animation: popdot 3s linear -2s infinite;margin-left:calc(50% - "+(((settings.radius/4)/2)*4)+"px);margin-top: calc(50% - "+(((settings.radius/4)/2))+"px);";

				var Dot1= $('<svg height="'+((settings.radius*2))+'" width="'+((settings.radius*2))+'"><circle style="transform-origin: '+settings.radius+'px '+settings.radius+'px; animation: popdot 1.5s linear 0s infinite" cx="50%" cy="50%" r="'+(settings.radius/2)+'" fill="rgba(0,0,0,0)" stroke="'+settings.colors[0]+'"></circle></svg>');

				$(this).append(Dot1);
				break;

			case 'hordots':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: spreaddot1-hor 1s linear 0s infinite;margin-left:"+settings.radius+"px;margin-top:"+settings.radius+"px;";
				var Dot2Style = "animation: spreaddot2-hor 1s linear 0.5s infinite;margin-left:"+settings.radius+"px;margin-top:"+settings.radius+"px;";
				var Dot3Style = "animation: spreaddot3-hor 1s linear 0s infinite;margin-left:"+settings.radius+"px;margin-top:"+settings.radius+"px;z-index:100";
				var Dot4Style = "animation: spreaddot4-hor 1s linear 0.5s infinite;margin-left:"+settings.radius+"px;margin-top:"+settings.radius+"px;";
				var Dot5Style = "animation: spreaddot5-hor 1s linear 0s infinite;margin-left:"+settings.radius+"px;margin-top:"+settings.radius+"px;";


				var Dot1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[3]+'"></circle></svg>');
				var Dot5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[4]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(Dot5);
				break;

			case 'verdots':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: spreaddot1-ver 1s linear 0s infinite;margin-left:calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot2Style = "animation: spreaddot2-ver 1s linear 0.5s infinite;margin-left:calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot3Style = "animation: spreaddot3-ver 1s linear 0s infinite;margin-left:calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;z-index:100";
				var Dot4Style = "animation: spreaddot4-ver 1s linear 0.5s infinite;margin-left:calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot5Style = "animation: spreaddot5-ver 1s linear 0s infinite;margin-left:calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";


				var Dot1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[3]+'"></circle></svg>');
				var Dot5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[4]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(Dot5);
				break;

			case 'spreaddots':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: spreaddot1-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot2Style = "animation: spreaddot2-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot3Style = "animation: spreaddot3-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;z-index:100";
				var Dot4Style = "animation: spreaddot4-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Dot5Style = "animation: spreaddot5-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";


				var Dot1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[3]+'"></circle></svg>');
				var Dot5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="'+settings.colors[4]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(Dot5);
				break;

			case 'trailedspreaddots':

				var RGB1 = 'rgb(0,0,255,0.9)';
				var RGB2 = 'rgb(0,0,255,0.9)';
				var RGB3 = 'rgb(0,0,255,0.9)';
				var RGB4 = 'rgb(0,0,255,0.9)';
				// var RGB1 = HexToRGB(settings.colors[0]);
				// var RGB2 = HexToRGB(settings.colors[1]);
				// var RGB3 = HexToRGB(settings.colors[4]);
				// var RGB4 = HexToRGB(settings.colors[3]);
				console.log(RGB1);
				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: spreaddot1-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Rect1Style = "animation: trailedspreadrect1-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;transform: rotateZ(45deg);z-index:-100;";
				var Dot2Style = "animation: spreaddot2-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Rect2Style = "animation: trailedspreadrect2-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;transform: rotateZ(45deg);z-index:-100;";
				var Dot3Style = "animation: dspreaddot3-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;z-index:100";
				var Dot4Style = "animation: spreaddot4-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Rect4Style = "animation: trailedspreadrect4-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;transform: rotateZ(45deg);z-index:-100;";
				var Dot5Style = "animation: spreaddot5-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;";
				var Rect5Style = "animation: trailedspreadrect5-all 1s linear 0s infinite;margin-left: calc(50% - "+((settings.radius/12))+"px);margin-top:"+settings.radius+"px;transform: rotateZ(45deg);z-index:-100;";


				var Dot1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="rgba(200,100,0,1)"></circle></svg>');
				var Rect1= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Rect1Style+'"><defs><linearGradient id="grad1" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" style="stop-color:'+RGB1+';stop-opacity:1" /><stop offset="100%" style="stop-color:rgb(255,255,255);stop-opacity:0" /></linearGradient></defs><rect x="0" y="0" width="'+(((settings.radius/6)+1)/2)+'" height="'+((settings.radius/6)+1)+'" fill="url(#grad1)"			/></svg>');
				var Dot2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="rgba(200,100,0,1)"></circle></svg>');
				var Rect2= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Rect2Style+'"><defs><linearGradient id="grad2" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" style="stop-color:'+RGB2+';stop-opacity:1" /><stop offset="100%" style="stop-color:rgb(255,255,255);stop-opacity:0" /></linearGradient></defs><rect x="0" y="0" width="'+(((settings.radius/6)+1)/2)+'" height="'+((settings.radius/6)+1)+'" fill="url(#grad2)"			/></svg>');
				var Dot3= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="rgba(0,0,0,1)"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="rgba(200,100,0,1)"></circle></svg>');
				var Rect4= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Rect4Style+'"><defs><linearGradient id="grad3" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" style="stop-color:'+RGB3+';stop-opacity:1" /><stop offset="100%" style="stop-color:rgb(255,255,255);stop-opacity:0" /></linearGradient></defs><rect x="0" y="0" width="'+(((settings.radius/6)+1)/2)+'" height="'+((settings.radius/6)+1)+'" fill="url(#grad3)"			/></svg>');
				var Dot5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/12)+'" fill="rgba(200,100,0,1)"></circle></svg>');
				var Rect5= $('<svg height="'+((settings.radius/6)+1)+'" width="'+((settings.radius/6)+1)+'" style="'+Rect5Style+'"><defs><linearGradient id="grad4" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" style="stop-color:'+RGB4+';stop-opacity:1" /><stop offset="100%" style="stop-color:rgb(255,255,255);stop-opacity:0" /></linearGradient></defs><rect x="0" y="0" width="'+(((settings.radius/6)+1)/2)+'" height="'+((settings.radius/6)+1)+'" fill="url(#grad4)"			/></svg>');
				$(this).append(Dot1).append(Rect1).append(Dot2).append(Rect2).append(Dot3).append(Dot4).append(Rect4).append(Dot5).append(Rect5);
				break;

			case 'circledots':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden").css("animation", "circledotdiv 3s linear 0s infinite");

				//Waves CSS
				var Dot1Style = "animation: circledot1 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot2Style = "animation: circledot2 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot3Style = "animation: circledot3 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;z-index:100";
				var Dot4Style = "animation: circledot4 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot5Style = "animation: circledot5 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot6Style = "animation: circledot6 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot7Style = "animation: circledot7 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot8Style = "animation: circledot8 1.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";


				var Dot1= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot1Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot2= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot2Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot3= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot3Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot4= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot4Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[3]+'"></circle></svg>');
				var Dot5= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot5Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[0]+'"></circle></svg>');
				var Dot6= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot6Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[1]+'"></circle></svg>');
				var Dot7= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot7Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[2]+'"></circle></svg>');
				var Dot8= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot8Style+'"><circle cx="50%" cy="50%" r="'+(settings.radius/16)+'" fill="'+settings.colors[3]+'"></circle></svg>');

				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(Dot5).append(Dot6).append(Dot7).append(Dot8);
				break;

			case 'squares':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Dot1Style = "animation: squares 4.5s linear 0s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;";
				var Dot2Style = "animation: squares 4.5s linear 0.5s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;";
				var Dot3Style = "animation: squares 4.5s linear 1s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;z-index:100";
				var Dot4Style = "animation: squares 4.5s linear 1.5s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var DotCenterStyle = "animation: squares 4.5s linear 2s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot6Style = "animation: squares 4.5s linear 2.5s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot7Style = "animation: squares 4.5s linear 3s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";
				var Dot8Style = "animation: squares 4.5s linear 3.5s infinite;margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";
				var Dot9Style = "animation: squares 4.5s linear 4s infinite;margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";


				var Dot1= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot1Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot2= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot2Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot3= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot3Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot4= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot4Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var DotCenter= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+DotCenterStyle+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[3]+';" /></svg>');
				var Dot6= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot6Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var Dot7= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot7Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot8= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot8Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot9= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot9Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				$(this).append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(DotCenter).append(Dot6).append(Dot7).append(Dot8).append(Dot9);
				break;

			case '3Dsquares':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden").css('animation','threedsqdiv 1s linear 0s infinite');

				//Waves CSS
				var Dot1Style = "margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;";
				var Dot2Style = "margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;";
				var Dot3Style = "margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*4)+"px;z-index:100";
				var Dot4Style = "margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var DotCenterStyle = "margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot6Style = "margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16))+"px;";
				var Dot7Style = "margin-left:"+((settings.radius)-(settings.radius/16)*4)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";
				var Dot8Style = "margin-left:"+((settings.radius)-(settings.radius/16))+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";
				var Dot9Style = "margin-left:"+((settings.radius)-(settings.radius/16)*-2)+"px;margin-top:"+((settings.radius)-(settings.radius/16)*-2)+"px;";

				var jRoll3DSquareFace1 = $('<div>', {id: 'jRoll3DSquareFace1'});
				var jRoll3DSquareFace2 = $('<div>', {id: 'jRoll3DSquareFace2'});
				$(this).append(jRoll3DSquareFace1).append(jRoll3DSquareFace2);
				$('#jRoll3DSquareFace1').css("animation", "threedsqf1 1s linear 0s infinite").css("transform","perspective(200px)").css('transform-origin','50% 50% 0px');
				$('#jRoll3DSquareFace2').css('animation', 'threedsqf2 1s linear 0s infinite').css("transform","perspective(200px)").css('transform-origin','50% '+(settings.radius+((settings.radius/8)*2))+'px 0px');
				var Dot1= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot1Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot2= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot2Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot3= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot3Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot4= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot4Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var DotCenter= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+DotCenterStyle+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[3]+';" /></svg>');
				var Dot6= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot6Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var Dot7= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot7Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot8= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot8Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot9= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot9Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');

				var Dot13D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot1Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot23D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot2Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot33D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot3Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot43D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot4Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var DotCenter3D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+DotCenterStyle+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[3]+';" /></svg>');
				var Dot63D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot6Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var Dot73D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot7Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				var Dot83D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot8Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Dot93D= $('<svg height="'+((settings.radius/8))+'" width="'+((settings.radius/8))+'" style="'+Dot9Style+'"><rect width="'+((settings.radius/8))+'" height="'+((settings.radius/8))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				$('#jRoll3DSquareFace1').append(Dot1).append(Dot2).append(Dot3).append(Dot4).append(DotCenter).append(Dot6).append(Dot7).append(Dot8).append(Dot9);
				$('#jRoll3DSquareFace2').append(Dot13D).append(Dot23D).append(Dot33D).append(Dot43D).append(DotCenter3D).append(Dot63D).append(Dot73D).append(Dot83D).append(Dot93D);
				break;

			case 'stackedsquares':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var Sq1Style = "animation: stackedsquare 1.5s linear 0s infinite;transform: perspective(500px)rotateX(90deg)rotateZ(-45deg)translateZ(0px);position:absolute;opacity:0;left:25%;";
				var Sq2Style = "animation: stackedsquare 1.5s linear 0.5s infinite;transform: perspective(500px)rotateX(90deg)rotateZ(-45deg)translateZ(-50px);position:absolute;opacity:0;left:25%;";
				var Sq3Style = "animation: stackedsquare 1.5s linear 1s infinite;transform: perspective(500px)rotateX(90deg)rotateZ(-45deg)translateZ(-100px);position:absolute;opacity:0;left:25%;";

				var Sq1= $('<svg height="'+((settings.radius))+'" width="'+((settings.radius))+'" style="'+Sq1Style+'"><rect width="'+((settings.radius))+'" height="'+((settings.radius))+'" style="fill:'+settings.colors[2]+';" /></svg>');
				var Sq2= $('<svg height="'+((settings.radius))+'" width="'+((settings.radius))+'" style="'+Sq2Style+'"><rect width="'+((settings.radius))+'" height="'+((settings.radius))+'" style="fill:'+settings.colors[1]+';" /></svg>');
				var Sq3= $('<svg height="'+((settings.radius))+'" width="'+((settings.radius))+'" style="'+Sq3Style+'"><rect width="'+((settings.radius))+'" height="'+((settings.radius))+'" style="fill:'+settings.colors[0]+';" /></svg>');
				$(this).append(Sq1).append(Sq2).append(Sq3);
				break;

			case 'waterdrop':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var waterdropStyle = "animation: waterdropdrop 1.5s linear 0s infinite;";
				var waterwaveStyle = "transform:perspective(500px)rotateX(50deg)scale(1,1);animation: waterdropwave 3s linear 0s infinite;";
				var waterwaveStyleIn = "transform:perspective(500px)rotateX(50deg)scale(1,1);animation: waterdropwaveIn 3s linear 0s infinite;";
				var waterwaveStyle2 = "transform:perspective(500px)rotateX(50deg)scale(1,1);opacity:0;animation: waterdropwave 3s linear 1.5s infinite;";
				var waterwaveStyle2In = "transform:perspective(500px)rotateX(50deg)scale(1,1);opacity:0;animation: waterdropwaveIn 3s linear 1.5s infinite;";


				var waterdrop= $('<svg height="'+((settings.radius)/4)+'" width="'+((settings.radius)*2)+'" viewBox="518.234 280.146 200 200" style="'+waterdropStyle+'"><path fill="'+settings.colors[0]+'" stroke="'+settings.colors[3]+'" stroke-width="10" stroke-miterlimit="10" d="M618.235,468.333c-37.982,0-45.104-29.489-45.104-29.489c-10.739-48.237,45.29-145.826,45.29-145.826s54.219,93.03,45.29,145.826C663.712,438.844,656.217,468.333,618.235,468.333z"/></svg>');
				var outRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollOutRing" style="'+waterwaveStyle+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+((settings.radius/2)-0)+'" stroke="'+settings.colors[2]+'" stroke-width="2" fill="transparent"></circle></svg>');
				var InRing= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollInRing" style="'+waterwaveStyleIn+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+((settings.radius/4)-0)+'" stroke="'+settings.colors[2]+'" stroke-width="1" fill="transparent"></circle></svg>');
				var outRing2= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollOutRing" style="'+waterwaveStyle2+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+((settings.radius/2)-0)+'" stroke="'+settings.colors[2]+'" stroke-width="2" fill="transparent"></circle></svg>');
				var InRing2= $('<svg height="'+(settings.radius*2)+'" width="'+(settings.radius*2)+'" id="jRollInRing" style="'+waterwaveStyle2In+'"><circle cx="'+settings.radius+'" cy="'+settings.radius+'" r="'+((settings.radius/4)-0)+'" stroke="'+settings.colors[2]+'" stroke-width="1" fill="transparent"></circle></svg>');

				$(this).append(waterdrop).append(outRing).append(InRing).append(outRing2).append(InRing2);
				break;

			case 'eq':

				//Parent CSS
				$(this).css("width", settings.radius*2+'px').css("height",settings.radius*2+'px').css("overflow","hidden");

				//Waves CSS
				var bar1Style = "animation: eq 2s linear -1s infinite;margin-left:calc(50% - "+(((settings.radius)/12)*2.5)+"px);top:calc(50% - "+((settings.radius)/4)+"px );";
				var bar2Style = "animation: eq 2s linear 0s infinite;margin-left:calc(50% - "+(((settings.radius)/12)*1.25)+"px);top:calc(50% - "+((settings.radius)/4)+"px );";
				var bar3Style = "animation: eq 2s linear -0.8s infinite;margin-left:50%;top:calc(50% - "+((settings.radius)/4)+"px );";
				var bar4Style = "animation: eq 2s linear -1.2s infinite;margin-left:calc(50% + "+(((settings.radius)/12)*1.25)+"px);top:calc(50% - "+((settings.radius)/4)+"px );";
				var bar5Style = "animation: eq 2s linear -0.4s infinite;margin-left:calc(50% + "+(((settings.radius)/12)*2.5)+"px);top:calc(50% - "+((settings.radius)/4)+"px );";

				var bar1= $('<svg height="'+((settings.radius)/2)+'" width="'+((settings.radius)*2)+'" style="'+bar1Style+'"><rect x="0" y="0" rx="4" ry="4" width="'+((settings.radius)/12)+'" height="'+((settings.radius)/2)+'" style="fill:'+settings.colors[3]+';opacity:0.5" /></svg>');
				var bar2= $('<svg height="'+((settings.radius)/2)+'" width="'+((settings.radius)*2)+'" style="'+bar2Style+'"><rect x="0" y="0" rx="4" ry="4" width="'+((settings.radius)/12)+'" height="'+((settings.radius)/2)+'" style="fill:'+settings.colors[1]+';opacity:0.5" /></svg>');
				var bar3= $('<svg height="'+((settings.radius)/2)+'" width="'+((settings.radius)*2)+'" style="'+bar3Style+'"><rect x="0" y="0" rx="4" ry="4" width="'+((settings.radius)/12)+'" height="'+((settings.radius)/2)+'" style="fill:'+settings.colors[0]+';opacity:0.5" /></svg>');
				var bar4= $('<svg height="'+((settings.radius)/2)+'" width="'+((settings.radius)*2)+'" style="'+bar4Style+'"><rect x="0" y="0" rx="4" ry="4" width="'+((settings.radius)/12)+'" height="'+((settings.radius)/2)+'" style="fill:'+settings.colors[2]+';opacity:0.5" /></svg>');
				var bar5= $('<svg height="'+((settings.radius)/2)+'" width="'+((settings.radius)*2)+'" style="'+bar5Style+'"><rect x="0" y="0" rx="4" ry="4" width="'+((settings.radius)/12)+'" height="'+((settings.radius)/2)+'" style="fill:'+settings.colors[4]+';opacity:0.5" /></svg>');

				$(this).append(bar1).append(bar2).append(bar3).append(bar4).append(bar5);
				break;

		}

    }

}(jQuery));

function HexToRGB(hex){
				var patt = /^#([\da-fA-F]{2})([\da-fA-F]{2})([\da-fA-F]{2})$/;
				var matches = patt.exec(hex);
				var rgb = "rgb("+parseInt(matches[1], 16)+","+parseInt(matches[2], 16)+","+parseInt(matches[3], 16)+");";
				return rgb;
}
